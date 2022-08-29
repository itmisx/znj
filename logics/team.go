package logics

import (
	"time"

	"znj/models"

	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// Team 团队操作
type Team struct{}

// List 团队列表
func (t Team) List(db *gorm.DB) (list []models.Teams, err error) {
	var count int64
	db.Find(&list).Count(&count)

	return
}

// Save 保存（含新增和编辑）
//
// @param team     待保存团队数据
// @param adminID  操作人id
func (t Team) Save(db *gorm.DB,
	team models.Teams,
	adminID string,
) (err error) {
	var (
		teamOld      models.Teams
		teamMember   models.TeamMember
		primaryKeyID string = ""
		rowsAffected int64  = 0
	)

	// 获取操作人id
	team.CreateID = adminID
	// 设置更新时间
	team.UpdateTime = time.Now().Unix()
	// 编辑
	if team.ID != "" {

		// 如果为编辑，先判断id的合法性
		rowsAffected = db.Table("teams").
			Where("id = ?", team.ID).
			First(&teamOld).
			RowsAffected
		if rowsAffected != 1 {
			// 要编辑的团队不存在
			err = errorx.New("", 202003)
			return
		}

		// 判断是否具有编辑的权限（条件：属于该团队，且未管理员）
		rowsAffected = db.Model(&teamMember).
			Where("team_id = ? and user_id = ? and (role = 1 or role = 2)", team.ID, adminID).
			First(&teamMember).
			RowsAffected

		// 没有权限
		if rowsAffected != 1 {
			err = errorx.New("", 202004)
			return
		}

		if team.Name != "" {
			teamOld.Name = team.Name
		}

		teamOld.UpdateTime = time.Now().Unix()

		// 更新内容
		rowsAffected = db.Model(&models.Teams{}).
			Where("id = ?", teamOld.ID).
			Updates(&teamOld).
			RowsAffected
		if rowsAffected != 1 {
			err = errorx.New("", 202002)
			return
		}
		return

	}
	// 新增
	if team.Name == "" {
		err = errorx.New("", 202005)
		return
	}
	// 设置主键id
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	team.ID = primaryKeyID
	// 设置创建时间
	team.CreateTime = time.Now().Unix()
	// 新增团队
	rowsAffected = db.Create(&team).RowsAffected
	if rowsAffected < 1 {
		err = errorx.New("", 202001)
		return
	}
	// 增加团队成员-管理员
	rowsAffected = t.CreateTeamMember(db, team.ID, adminID, 1)
	if rowsAffected < 1 {
		err = errorx.New("", 202001)
		return
	}
	return
}

// CreateTeamMember 添加团队成员
func (t Team) CreateTeamMember(db *gorm.DB,
	teamID string,
	userID string,
	role int64,
) (rowsAffected int64) {
	var teamMember models.TeamMember

	// 设置主键id
	primaryKeyID, err := helper.UUID{}.String()
	if err != nil {
		return
	}
	teamMember.ID = primaryKeyID
	teamMember.TeamID = teamID

	teamMember.UserID = userID

	teamMember.Role = role

	teamMember.CreateTime = time.Now().Unix()

	// 增加团队成员
	rowsAffected = db.Create(&teamMember).RowsAffected

	return rowsAffected
}

// MemberList 获取成员团队成员列表
// projectId 项目id adminID 管理员id
func (t Team) MemberList(db *gorm.DB,
	teamID string,
	projectID string,
	adminID string,
) (list interface{}, userRole int64, err error) {
	type MemberList struct {
		User string `json:"user"`
		Role int64  `json:"role"`
		Name string `json:"name"`
		ID   string `json:"id"`
		Team string `json:"team"`
	}

	var (
		memberList    []MemberList
		projectRecord models.Projects
	)
	// 当前用户的团队角色
	userRole = 3

	// 获取项目所属团队
	if teamID == "" {
		db.Table("projects").Where("id = ?", projectID).First(&projectRecord)
		teamID = projectRecord.TeamID
	}
	// 判断当前用户是否具有访问权限
	userRole, err = t.HasAccessAuthority(db, teamID, adminID)

	if err != nil {
		return
	}

	// 获取成员列表
	db.Table("team_member tm").
		Select("tm.user_id as id,tm.role,teams.name as team,users.name as name,users.user").
		Joins("join users on users.id = tm.user_id").
		Joins("join teams on teams.id = tm.team_id").
		Where("tm.team_id = ?", teamID).
		Find(&memberList)
	list = memberList

	return
}

// AddTeamMember 添加团队成员(含新增和邀请)
// 团队id必须（要加入的团队）
// 如果有用户id则为新增，否则为邀请
func (t Team) AddTeamMember(db *gorm.DB,
	user models.Users,
	teamID string,
	adminID string,
) (err error) {
	var (
		userID       string
		rowsAffected int64  = 0
		primaryKeyID string = ""
		teamMember   models.TeamMember
		userlogic    = User{}
	)

	if teamID == "" {
		err = errorx.New("所属团队不能为空", 202101)
		return
	}

	// 验证所属团队是否存在
	db.Table("teams").
		Where("id=?", teamID).Count(&rowsAffected)

	if rowsAffected <= 0 {
		err = errorx.New("所属团队不存在", 202102)
		return
	}

	// 判断操作人的管理员权限
	db.Table("team_member").
		Where("team_id = ? and user_id = ? and (role = 1 or role = 2)", teamID, adminID).
		Count(&rowsAffected)

	if rowsAffected <= 0 {
		err = errorx.New("非管理员，无权限", 202004)
		return
	}

	// 添加成员 - 后面可能会增加邀请成员
	if user.ID == "" {
		userID, err = userlogic.CreateUser(db, user, teamID)
		if err != nil || userID == "" {
			return
		}
		user.ID = userID
	}
	// 判断userid是否合法
	db.Table("users").
		Where("id=?", user.ID).
		Count(&rowsAffected)
	if rowsAffected <= 0 {
		err = errorx.New("", 202206)
		return
	}

	// 添加成员到团队
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	teamMember.ID = primaryKeyID
	teamMember.TeamID = teamID
	teamMember.UserID = user.ID
	teamMember.Role = 3 // 默认为普通用户 1.超级管理员 2.管理员
	teamMember.CreateTime = time.Now().Unix()
	rowsAffected = db.Model(&teamMember).
		Create(&teamMember).RowsAffected

	if rowsAffected <= 0 {
		err = errorx.New("", 202103)
		return
	}

	return
}

// SetWorkTeam 设置当前的工作团队
//
// @param workTeamID 团队
// @Param adminID    操作人id
func (t Team) SetWorkTeam(db *gorm.DB,
	workTeamID string,
	adminID string,
) (err error) {
	var rowsAffected int64 = 0

	// 保存teamID
	if workTeamID != "" {

		// 保存teamID
		rowsAffected = db.Table("users").
			Where("id = ?", adminID).
			Update("work_team_id", workTeamID).RowsAffected
		if rowsAffected > 0 {
			return
		}
	}
	err = errorx.New("", 202207)
	return
}

// GetWorkTeam 获取当前的工作团队
func (t Team) GetWorkTeam(db *gorm.DB,
	adminID string,
) (teamID string) {
	var (
		count int64
		user  models.Users
	)

	// 获取teamID
	db.Table("users").
		Where("id = ?", adminID).
		First(&user).Count(&count)

	if count == 1 {
		teamID = user.WorkTeamID
		return
	}

	return
}

// ResetPassword 重置密码
//
// @param userID 成员id
// @param teamID 团队id
// @param adminID  操作人id
func (t Team) ResetPassword(db *gorm.DB,
	userID string,
	password string,
	adminID string,
) (err error) {
	var (
		count        int64
		rowsAffected int64
		userOld      models.Users
	)
	// 判断用户id是否存在
	db.Table("users").
		Where("id = ?", userID).
		Take(&userOld).
		Count(&count)
	if count <= 0 {
		err = errorx.New("", 202206)
		return
	}

	// 判断是否为管理员
	db.Table("team_member").
		Where("user_id = ?", adminID).
		Where("team_id = ? and (role = 1 or role = 2)", userOld.TeamID).
		Count(&count)
	if count <= 0 {
		err = errorx.New("非管理员，无权限", 202004)
		return
	}

	userOld.Password = helper.Md5(helper.Md5(password))
	// 重置密码
	rowsAffected = db.Model(&models.Users{}).
		Where("id = ?", userID).
		Updates(&userOld).
		RowsAffected

	if rowsAffected != 1 {
		err = errorx.New("", 202104)
		return
	}
	return
}

// SetAdminRole 设为管理员
//
// @param userID 成员id
// @param teamID 团队id
// @param adminID 操作人id
func (t Team) SetAdminRole(db *gorm.DB,
	userID string,
	teamID string,
	adminID string,
) (err error) {
	var (
		count        int64
		rowsAffected int64
		teamMember   models.TeamMember
	)

	if teamID == "" {
		err = errorx.New("所属团队不能为空", 202101)
		return
	}

	// 判断团队是否存在
	db.Table("teams").
		Where("id = ?", teamID).
		Count(&count)
	if count <= 0 {
		err = errorx.New("所属团队不存在", 202102)
		return
	}

	// 判断用户是否属于该团队
	db.Table("team_member").
		Where("team_id = ? and user_id = ?", teamID, userID).
		Take(&teamMember).
		Count(&count)
	if count <= 0 {
		err = errorx.New("", 202007)
		return
	}

	// 判断是否为管理员
	db.Table("team_member").
		Where("user_id = ?", adminID).
		Where("team_id = ? and role = 1", teamID).
		Count(&count)
	if count <= 0 {
		err = errorx.New("", 202004)
		return // 非管理员，无权限
	}

	// 修改角色
	teamMember.Role = 2
	rowsAffected = db.Model(&models.TeamMember{}).
		Where("team_id = ? and user_id = ?", teamID, userID).
		Updates(&teamMember).
		RowsAffected
	if rowsAffected != 1 {
		err = errorx.New("", 202105)
		return
	}

	return
}

// HasAccessAuthority 验证是否具有团队访问权限
func (t Team) HasAccessAuthority(db *gorm.DB, teamID string, userID string) (userRole int64, err error) {
	var count int64
	// 默认角色为项目成员
	userRole = 3
	// 判断是否为团队创建人
	db.Table("teams").
		Where("id = ? and create_id = ?", teamID, userID).
		Count(&count)
	if count > 0 {
		userRole = 1
		return
	}

	// 判断是否为团队管理员
	db.Table("team_member").
		Where("team_id = ? and user_id = ? and role = 2", teamID, userID).
		Count(&count)
	if count > 0 {
		userRole = 2
		return
	}
	// 判断是否为项目成员
	db.Table("team_member").
		Where("team_id = ? and user_id = ? and role = 3", teamID, userID).
		Count(&count)
	if count > 0 {
		userRole = 3
		return
	}
	// 不属于团队成员，无权限
	err = errorx.New("", 202006)
	return
}
