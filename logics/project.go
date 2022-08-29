package logics

import (
	"strings"
	"time"

	"znj/models"

	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// Project 项目操作
type Project struct{}

// List 项目列表
//
// @param workTeamID 工作团队id
// @param adminID    操作人id
func (p Project) List(db *gorm.DB,
	workTeamID string,
	adminID string,
) (list []models.Projects, err error) {
	// 通过团队id，查询团队下的项目
	db.Table("projects p").
		Select("p.*").
		Joins("join project_member pm on pm.project_id = p.id").
		Where("pm.user_id = ?", adminID).
		Where("p.team_id = ?", workTeamID).
		Order("star_time desc").
		Order("create_time asc").
		Find(&list)

	return
}

// Save 项目编辑
func (p Project) Save(db *gorm.DB,
	teamID string,
	project models.Projects,
	adminID string,
) (err error) {
	var (
		projectOld    models.Projects
		projectMember models.ProjectMember
		primaryKeyID  string
		rowsAffected  int64 = 0
	)

	// 设置更新时间
	project.UpdateTime = time.Now().Unix()

	// 编辑
	if project.ID != "" {

		// 如果为编辑，先判断id的合法性
		db.Table("projects").Where("id = ?", project.ID).First(&projectOld).Count(&rowsAffected)
		if rowsAffected != 1 {
			err = errorx.New("要编辑的项目不存在", 204003)
			return
		}

		// 判断是否具有编辑的权限（条件：属于该团队，且未管理员）
		rowsAffected = db.Table("project_member").
			Where("project_id = ? && user_id = ? ", project.ID, adminID).
			First(&projectMember).
			RowsAffected

		// 没有权限
		if rowsAffected != 1 {
			err = errorx.New("不属于项目成员，无权限", 204004)
			return
		}

		// 项目名称
		if project.Name != "" {
			projectOld.Name = project.Name
		}

		// 项目描述
		if project.Desc != "" {
			projectOld.Desc = project.Desc
		}

		// 置顶
		if project.StarTime > 0 {
			projectOld.StarTime = time.Now().Unix()
		}
		projectOld.UpdateTime = time.Now().Unix()
		// 更新内容
		rowsAffected = db.Model(&models.Projects{}).
			Where("id = ?", projectOld.ID).
			Updates(&projectOld).
			RowsAffected
		if rowsAffected != 1 {
			err = errorx.New("", 204002)
			return
		}

		return

	}
	/*新增*/

	// 主键id
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	project.ID = primaryKeyID

	if teamID == "" {
		err = errorx.New("", 204007)
		return
	}
	project.TeamID = teamID
	// 项目名称
	if project.Name == "" {
		err = errorx.New("项目名称不能为空", 204005)
		return
	}
	// 项目描述
	if project.Desc == "" {
		err = errorx.New("项目描述不能为空", 204006)
		return
	}
	// 项目创建人
	project.CreateID = adminID
	// 设置创建时间
	project.CreateTime = time.Now().Unix()
	project.UpdateTime = time.Now().Unix()

	// 新增项目
	rowsAffected = db.Create(&project).RowsAffected
	if rowsAffected < 1 {
		err = errorx.New("新增项目失败", 204001)
		return
	}
	// 增加项目成员
	rowsAffected, _ = p.CreateProjectMember(db, project.ID, adminID)
	if rowsAffected < 1 {
		err = errorx.New("新增项目失败", 204001)
		return
	}
	return
}

// CreateProjectMember 创建团队成员
func (p Project) CreateProjectMember(db *gorm.DB,
	ProjectID string,
	userID string,
) (rowsAffected int64, err error) {
	var projectMember models.ProjectMember
	// 设置主键id
	primaryKeyID, err := helper.UUID{}.String()
	if err != nil {
		return
	}
	projectMember.ID = primaryKeyID
	// 项目id
	projectMember.ProjectID = ProjectID
	// 用户id
	projectMember.UserID = userID
	// 创建时间
	projectMember.CreateTime = time.Now().Unix()

	// 增加团队成员
	rowsAffected = db.Create(&projectMember).RowsAffected

	return
}

// TodoTypeSave 待办类型保存(含新增和编辑)
//
// @param todoType 待办类型
// @param adminID  操作人id
func (p Project) TodoTypeSave(db *gorm.DB,
	todoType models.TodoType,
	adminID string,
) (err error) {
	var (
		todoTypeOld  = models.TodoType{}
		primaryKeyID string
	)

	// 编辑操作
	if todoType.ID != "" {
		var count int64
		// 判断是否存在
		db.Table("todo_type").Where("id = ?", todoType.ID).Find(&todoTypeOld).Count(&count)
		if count <= 0 {
			err = errorx.New("", 204103)
			return
		}

		// 判断项目权限
		err = p.HasAccessAuthority(db, todoTypeOld.ProjectID, adminID)
		if err != nil {
			return
		}

		// 类型名字
		if todoType.Name != "" {
			todoTypeOld.Name = todoType.Name
		}
		// 置顶时间
		if todoType.StarTime > 0 {
			todoTypeOld.StarTime = time.Now().Unix()
		}

		todoTypeOld.UpdateTime = time.Now().Unix()
		// 更新
		rowsAffected := db.Model(&models.TodoType{}).
			Where("id = ?", todoType.ID).
			Updates(&todoTypeOld).
			RowsAffected

		if rowsAffected != 1 {
			err = errorx.New("", 204105)
		}
		return
	}

	/*新增操作*/

	// 所属项目不能为空
	if todoType.ProjectID == "" {
		err = errorx.New("", 204101)
		return
	}

	// 判断项目权限
	err = p.HasAccessAuthority(db, todoType.ProjectID, adminID)
	if err != nil {
		return
	}

	// 任务类型名称不能为空
	if todoType.Name == "" {
		err = errorx.New("", 204102)
		return
	}
	// 获取主键
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	todoType.ID = primaryKeyID
	// 排序
	if todoType.Sort == 0 {
		todoType.Sort = 9999
	}
	// 创建时间
	todoType.CreateTime = time.Now().Unix()
	// 更新时间
	todoType.UpdateTime = time.Now().Unix()

	// 保存
	rowsAffected := db.Model(&todoType).Create(todoType).RowsAffected
	if rowsAffected <= 0 {
		err = errorx.New("", 204104)
		return
	}

	return
}

// TodoTypeList 获取待办类型列表
//
// @param projectID 项目id
// @param adminID   操作人id
func (p Project) TodoTypeList(db *gorm.DB,
	projectID string,
	adminID string,
) (list []models.TodoType, err error) {
	var doc Doc

	// 所属项目不能为空
	if projectID == "" {
		err = errorx.New("", 204101)
		return
	}

	// 判断项目权限
	err = p.HasAccessAuthority(db, projectID, adminID)
	if err != nil {
		return
	}

	doc.SetLatestProjects(db, adminID, projectID)

	// 待办类型列表
	db.Table("todo_type").
		Where("project_id = ?", projectID).
		Order("sort asc").
		Order("create_time asc").
		Find(&list)

	return
}

// MemberSet 项目成员设置
func (p Project) MemberSet(db *gorm.DB,
	projectID string,
	IDs string,
	adminID string,
) (err error) {
	var (
		projectMember models.ProjectMember
		count         int64 = 0
		primaryKeyID  string
		memberIDSlice []string
	)

	// 所属项目不能为空
	if projectID == "" {
		err = errorx.New("", 204101)
		return
	}

	// 获取设置的id集合
	if IDs != "" {
		memberIDSlice = strings.Split(IDs, ",")
	}

	// 判断项目是否存在
	db.Table("projects").
		Where("id = ?", projectID).
		Count(&count)
	if count < 0 {
		err = errorx.New("", 204003)
		return
	}

	// 验证项目权限
	err = p.HasAccessAuthority(db, projectID, adminID)
	if err != nil {
		return
	}
	// 删除原有项目成员
	err = db.Where("project_id = ?", projectID).
		Delete(projectMember).Error
	if err != nil {
		err = errorx.New(err.Error(), 204008)
		return
	}

	// 迭代并组装成项目成员数组
	for _, IDStr := range memberIDSlice {
		ID := IDStr
		// 验证ID是否合法
		db.Table("users").
			Where("id = ?", ID).
			Count(&count)
		if count == 1 {
			// 获取主键
			primaryKeyID, err = helper.UUID{}.String()
			if err != nil {
				return
			}
			temp := models.ProjectMember{}
			temp.ID = primaryKeyID
			temp.ProjectID = projectID
			temp.UserID = ID
			temp.CreateTime = time.Now().Unix()

			// 插入成员
			err := db.Create(&temp).Error
			if err != nil {
				continue
			}
		}
	}
	return
}

// MemberList 项目成员列表
//
// @param projectID 项目id
// @param adminID   操作人id
func (p Project) MemberList(db *gorm.DB,
	projectID string,
	adminID string,
) (list interface{}, err error) {
	var memberList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	// 验证权限
	err = p.HasAccessAuthority(db, projectID, adminID)
	if err != nil {
		return
	}

	db.Table("project_member pm ").
		Joins("join users u on u.id = pm.user_id").
		Select("u.id,u.name").
		Where("pm.project_id = ?", projectID).
		Find(&memberList)

	list = memberList

	return
}

// TodoTypeBatchSort 待办类型批量排序
// todoTypeID-sort的字符串以逗号分隔的形式组合
func (p Project) TodoTypeBatchSort(db *gorm.DB, IDSorts string) (err error) {
	// 分隔要排序的待办类型
	IDSortsArray := strings.Split(IDSorts, ",")
	// 遍历获取每个待办类型的id和sort，并更新
	for _, idSort := range IDSortsArray {
		temp := strings.Split(idSort, "-")
		if len(temp) != 2 {
			continue
		}
		todoTypeID := temp[0]
		sort := temp[1]
		if todoTypeID == "" || sort == "" {
			continue
		}
		db.Table("todo_type").
			Where("id = ?", todoTypeID).
			Update("sort", sort)
	}

	return
}

// HasAccessAuthority 验证项目权限
func (p Project) HasAccessAuthority(db *gorm.DB,
	projectID string, userID string,
) (err error) {
	var (
		team    models.Teams
		project models.Projects
		count   int64
	)

	// 判断是否为项目成员
	db.Model(models.ProjectMember{}).
		Where("project_id = ? and user_id = ?", projectID, userID).
		Count(&count)
	if count > 0 {
		return
	}

	// 判断是否为项目管理员
	db.Model(project).Where("id = ?", projectID).First(&project)
	if project.CreateID == userID {
		return
	}
	// 判断是否为团队管理员
	if project.TeamID != "" {
		db.Model(team).
			Where("id = ? && user_id = ? && role = 1", project.TeamID, userID).
			Count(&count)
		if count > 0 {
			return
		}
	}
	// 不属于团队成员，无权限
	err = errorx.New("不属于团队成员，无权限", 204004)
	return
}
