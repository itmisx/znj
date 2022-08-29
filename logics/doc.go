package logics

import (
	"strings"
	"time"

	"znj/models"

	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// Doc 文档操作
type Doc struct{}

// ChapterTree 章节树结构
type ChapterTree struct {
	models.Chapters
	Expanded bool          `gorm:"-" json:"expanded"`
	Leaf     bool          `gorm:"-" json:"leaf"`
	Children []ChapterTree `gorm:"-" json:"children"`
}

// List 文档列表
//
// @param adminID 操作人id
// @param workTeamID  团队id
// @param projectID 项目id
//
// @return list 文档列表
// @return err  错误
func (d Doc) List(db *gorm.DB,
	adminID string,
	workTeamID string,
	projectID string,
) (list []models.Docs, err error) {
	// 团队文档
	if workTeamID != "" {
		// 确定团队访问权限
		t := Team{}
		_, err = t.HasAccessAuthority(db, workTeamID, adminID)
		if err != nil {
			return
		}
		// 获取团队文档列表
		db.Table("docs").
			Where("team_id = ?", workTeamID).
			Find(&list)
		// 项目文档
	} else if projectID != "" {
		// 判断项目访问权限
		p := Project{}
		err = p.HasAccessAuthority(db, projectID, adminID)
		if err != nil {
			return
		}
		// 获取项目文档列表
		db.Table("docs").
			Where("project_id = ?", projectID).
			Find(&list)
		// 个人文档
	} else {
		db.Table("docs").
			Where("user_id = ?", adminID).
			Find(&list)
	}

	return
}

// ChapterBatchSort 章节批量排序
//
// @param IDSorts  chapterID-sort的字符串以逗号分隔的形式组合
func (d Doc) ChapterBatchSort(db *gorm.DB, IDSorts string) (err error) {
	// 分隔要排序的待办类型
	IDSortsArray := strings.Split(IDSorts, ",")
	// 遍历获取每个待办类型的id和sort，并更新
	for _, idSort := range IDSortsArray {
		temp := strings.Split(idSort, "-")
		if len(temp) != 2 {
			continue
		}
		chapterID := temp[0]
		sort := temp[1]
		if chapterID == "" || sort == "" {
			continue
		}
		db.Table("chapters").
			Where("id = ?", chapterID).
			Update("sort", sort)
	}

	return
}

// Save 文档保存
//
// @param doc 文档模型
// @param adminID 操作人id
// @param workTeamID 工作团队
//
// @return error 错误
func (d Doc) Save(db *gorm.DB,
	doc models.Docs,
	adminID string,
	workTeamID string,
) (err error) {
	var (
		count        int64
		rowsAffected int64
		docOld       models.Docs
	)
	// 编辑文档
	if doc.ID != "" {

		// 确定文档是否存在
		db.Where("id = ?", doc.ID).Take(&docOld).Count(&count)
		if count <= 0 {
			err = errorx.New("", 205003)
			return
		}

		// 判断文档权限
		err = d.HasAccessAuthority(db, doc.ID, adminID)
		if err != nil {
			return
		}

		if doc.Name != "" {
			docOld.Name = doc.Name
		}

		docOld.UpdateTime = time.Now().Unix()
		rowsAffected = db.Model(&models.Docs{}).
			Where("id = ?", docOld.ID).
			Updates(&docOld).
			RowsAffected
		if rowsAffected != 1 {
			err = errorx.New("", 205002)
		}
		return
	}

	// 新增文档

	// 获取唯一主键
	primaryKeyID, err := helper.UUID{}.String()
	if err != nil {
		return
	}
	doc.ID = primaryKeyID
	// 文档名称不能为空
	if doc.Name == "" {
		err = errorx.New("", 205005)
		return
	}
	// 团队文档
	if doc.TeamID != "" {
		// 确定团队访问权限
		t := Team{}
		_, err = t.HasAccessAuthority(db, workTeamID, adminID)
		if err != nil {
			return
		}
		doc.TeamID = workTeamID
		// 项目文档
	} else if doc.ProjectID != "" {
		// 确定团队id是否存在
		db.Table("projects").
			Where("id = ?", doc.ProjectID).
			Count(&count)
		if count <= 0 {
			err = errorx.New("", 202102)
			return
		}
		// 确定项目访问权限
		p := Project{}
		err = p.HasAccessAuthority(db, doc.ProjectID, adminID)
		if err != nil {
			return
		}
		// 个人文档
	} else {
		doc.UserID = adminID
	}

	doc.CreateTime = time.Now().Unix()
	doc.UpdateTime = time.Now().Unix()

	rowsAffected = db.Create(doc).RowsAffected
	if rowsAffected != 1 {
		err = errorx.New("", 205001)
		return
	}
	return
}

// ChapterList 章节列表
//
// @param docID 文档id
// @param adminID 操作人id
//
// @return tree 文档章节树
func (d Doc) ChapterList(db *gorm.DB,
	docID string,
	adminID string,
) (tree interface{}, err error) {
	chapterTree := ChapterTree{}

	// 所属文档不能为空
	if docID == "" {
		err = errorx.New("", 205103)
		return
	}
	// 确认是否有文档访问权限
	err = d.HasAccessAuthority(db, docID, adminID)
	if err != nil {
		return
	}

	// 获取对应章节列表
	var list []ChapterTree
	db.Model(&models.Chapters{}).
		Order("sort asc,create_time asc").
		Where("doc_id = ? ", docID).
		Find(&list)

	// 初始化根节点
	chapterTree.ID = "0"
	chapterTree.PID = "0"
	chapterTree.Expanded = true
	chapterTree.Name = "目录"
	chapterTree.Leaf = false

	d.findChildChapter("0", &chapterTree, list)

	tree = []ChapterTree{chapterTree}

	return
}

// 递归查找子章节
func (d Doc) findChildChapter(pid string, chapterTree *ChapterTree, list []ChapterTree) {
	for _, v := range list {
		if v.PID == pid {
			// 整理一级章节
			chapterTree.Children = append(chapterTree.Children, v)
			// 递归生成章节树
			d.findChildChapter(v.ID, &(chapterTree.Children[len(chapterTree.Children)-1]), list)
			//
			if len(chapterTree.Children[len(chapterTree.Children)-1].Children) < 1 {
				chapterTree.Leaf = true
			} else {
				chapterTree.Leaf = false
			}
		}
	}
}

// ChapterSave 保存章节
//
// @param chapter 待保存的章节内容
// @param adminID 操作人id
//
// @return error 错误
func (d Doc) ChapterSave(db *gorm.DB,
	chapter models.Chapters,
	adminID string,
) (err error) {
	var (
		count        int64
		rowsAffected int64
		primaryKeyID string
		chapterOld   models.Chapters
	)

	/*编辑操作*/
	if chapter.ID != "" {

		// 确认章节是否存在
		db.Table("chapters").
			Where("id = ?", chapter.ID).
			First(&chapterOld).
			Count(&count)
		if count <= 0 {
			err = errorx.New("", 205104)
			return
		}
		// 确认文档操作权限
		err = d.HasAccessAuthority(db, chapterOld.DocID, adminID)
		if err != nil {
			return
		}
		// 章节名称
		if chapter.Name != "" {
			chapterOld.Name = chapter.Name
		}

		// 章节类型
		if chapter.DocType != 0 {
			chapterOld.DocType = chapter.DocType
		}

		// 所属章节
		if chapter.PID != "" && chapter.PID != chapterOld.PID {
			chapterOld.PID = chapter.PID
		}

		// 更新时间
		chapterOld.UpdateTime = time.Now().Unix()

		// 数据库更新
		rowsAffected = db.Model(&models.Chapters{}).
			Select("name,doc_type,pid,update_time").
			Where("id = ?", chapter.ID).
			Updates(map[string]interface{}{
				"name":        chapterOld.Name,
				"doc_type":    chapterOld.DocType,
				"pid":         chapterOld.PID,
				"update_time": chapterOld.UpdateTime,
			}).
			RowsAffected
		if rowsAffected != 1 {
			err = errorx.New("", 205102)
		}
		return
	}

	/*新增操作*/

	// 获取唯一主键
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	chapter.ID = primaryKeyID
	// 章节名称不能为空
	if chapter.Name == "" {
		err = errorx.New("", 205105)
		return
	}

	// 判断父级id
	if chapter.PID != "" && chapter.PID != "0" {
		// 确认章节是否存在
		db.Table("chapters").
			Where("id = ?", chapter.PID).
			First(&chapterOld).
			Count(&count)
		if count <= 0 {
			err = errorx.New("", 205104)
			return
		}
		// 确认文档操作权限
		err = d.HasAccessAuthority(db, chapterOld.DocID, adminID)
		if err != nil {
			return
		}
	} else {
		chapter.PID = "0"
	}
	// 所属文档,并确认文档访问权限
	if chapter.DocID == "" {
		err = errorx.New("", 205103)
		return
	}
	err = d.HasAccessAuthority(db, chapter.DocID, adminID)
	if err != nil {
		return
	}

	// 创建时间，更新时间
	chapter.CreateTime = time.Now().Unix()
	chapter.UpdateTime = time.Now().Unix()
	// 排序
	chapter.Sort = 9999

	// 创建人
	chapter.CreateID = adminID

	// 保存
	rowsAffected = db.Table("chapters").
		Create(chapter).RowsAffected
	if rowsAffected != 1 {
		err = errorx.New("", 205101)
	}

	return
}

// ChapterContent 获取章节内容
//
// @param chapterContent 章节内容
// @param adminID        操作人id
//
// @return content 章节内容
// @return err     错误
func (d Doc) ChapterContent(db *gorm.DB,
	chapterContent models.ChapterContent,
	adminID string,
) (content string, err error) {
	var (
		count   int64
		chapter models.Chapters
	)
	// 章节id不能为空
	if chapterContent.ChapterID == "" {
		err = errorx.New("", 205201)
		return
	}

	// 判断章节是否存在
	db.Table("chapters").
		Where("id = ?", chapterContent.ChapterID).
		First(&chapter).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 205104)
		return
	}
	// 判断章节访问权限
	err = d.HasAccessAuthority(db, chapter.DocID, adminID)
	if err != nil {
		return
	}

	// 保存至最近访问
	d.SetLatestChapters(db, chapterContent.ChapterID, adminID)

	// 获取章节内容
	db.Table("chapter_content").
		Where("version = ? and chapter_id = ?", chapter.Version, chapterContent.ChapterID).
		First(&chapterContent)

	content = chapterContent.Content
	return
}

// ChapterContentList 获取历史内容列表
func (d Doc) ChapterContentList(db *gorm.DB,
	chapterID string,
	adminID string,
) (list interface{}, err error) {
	var (
		count   int64
		chapter models.Chapters
	)

	type ChapterContentList struct {
		Version       string `json:"version"`
		Comment       string `json:"comment"`
		CreateTimeStr string `json:"create_time_str"`
		Name          string `json:"name"`
	}

	// 章节id不能为空
	if chapterID == "" {
		err = errorx.New("", 205201)
		return
	}

	// 判断章节是否存在
	db.Table("chapters").
		Where("id = ?", chapterID).
		Find(&chapter).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 205104)
		return
	}
	// 判断章节访问权限
	err = d.HasAccessAuthority(db, chapter.DocID, adminID)
	if err != nil {
		return
	}

	var chapterContentList []ChapterContentList

	// 获取章节内容
	db.Table("chapter_content as cc").
		Joins("left join users u on u.id = cc.create_id").
		Select("cc.version,cc.comment,FROM_UNIXTIME(cc.create_time) as create_time_str,u.name").
		Where("chapter_id = ?", chapterID).
		Order("cc.create_time desc").
		Find(&chapterContentList)

	list = chapterContentList
	return
}

// ChapterContentSave 保存章节内容
func (d Doc) ChapterContentSave(
	db *gorm.DB,
	chapterContent models.ChapterContent,
	userID string,
) (err error) {
	var (
		count        int64
		rowsAffected int64
		primaryKeyID string
		chapter      models.Chapters
	)
	// 章节id不能为空
	if chapterContent.ChapterID == "" {
		err = errorx.New("", 205201)
		return
	}

	// 判断章节是否存在
	db.Table("chapters").
		Where("id = ?", chapterContent.ChapterID).
		First(&chapter).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 205104)
		return
	}
	// 判断章节访问权限
	err = d.HasAccessAuthority(db, chapter.DocID, userID)
	if err != nil {
		return
	}
	// 获取唯一主键
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	chapterContent.Version = primaryKeyID
	chapterContent.CreateTime = time.Now().Unix()
	chapterContent.CreateID = userID

	// 新增章节内容
	rowsAffected = db.Table("chapter_content").Create(chapterContent).RowsAffected
	if rowsAffected != 1 {
		err = errorx.New("", 205201)
		return
	}

	// 设置章节内容版本
	chapter.Version = chapterContent.Version
	rowsAffected = db.Model(&models.Chapters{}).
		Where("id = ?", chapterContent.ChapterID).
		Updates(&chapter).
		RowsAffected
	if rowsAffected != 1 {
		err = errorx.New("", 205201)
	}

	return
}

// GetChapterRevision get章节版本内容
//
// @param chapterID 章节id
// @param chapterVersion 章节版本
// @param adminID 操作人id
func (d Doc) GetChapterRevision(
	db *gorm.DB,
	chapterID string,
	chapterVersion string,
	adminID string,
) (content string, err error) {
	var (
		chapterOld     models.Chapters
		chapterContent models.ChapterContent
	)

	// 章节id不能为空
	if chapterID == "" {
		err = errorx.New("", 205203)
		return
	}
	// 章节版本号不能为空
	if chapterVersion == "" {
		err = errorx.New("", 205204)
		return
	}

	// 验证章节是否存在
	db.Table("chapters").
		Where("id = ?", chapterID).
		First(&chapterOld)
	if chapterOld.ID == "" {
		err = errorx.New("", 205104)
		return
	}
	// 验证文档权限
	err = d.HasAccessAuthority(db, chapterOld.DocID, adminID)
	if err != nil {
		return
	}
	// 获取章节版本内容
	db.Table("chapter_content").
		Where("version = ?", chapterVersion).
		First(&chapterContent)

	content = chapterContent.Content
	return
}

// SetChapterRevision 回滚版本
//
// @param chapterID 章节id
// @param chapterVersion 章节版本
// @param adminID 操作人id
//
// @return err
func (d Doc) SetChapterRevision(
	db *gorm.DB,
	chapterID string,
	chapterVersion string,
	adminID string,
) (err error) {
	var (
		rowsAffected   int64
		chapterOld     models.Chapters
		chapterContent models.ChapterContent
	)

	// 所属章节不能为空
	if chapterID == "" {
		err = errorx.New("", 205203)
		return
	}
	// 章节版本不能空
	if chapterVersion == "" {
		err = errorx.New("", 205204)
		return
	}

	// 判断章节是否存在
	db.Table("chapters").
		Where("id = ?", chapterID).
		First(&chapterOld)
	if chapterOld.ID == "" {
		err = errorx.New("", 205104)
		return
	}
	// 判断版本是否存在
	db.Table("chapter_content").
		Where("version = ?", chapterVersion).
		First(&chapterContent)
	if chapterContent.Version == "" {
		err = errorx.New("", 205205)
		return
	}
	// 判断文档访问权限
	err = d.HasAccessAuthority(db, chapterOld.DocID, adminID)
	if err != nil {
		return
	}
	// 回滚版本
	chapterOld.Version = chapterVersion
	rowsAffected = db.Model(&models.Chapters{}).
		Where("id = ?", chapterID).
		Updates(&chapterOld).
		RowsAffected

	if rowsAffected != 1 {
		err = errorx.New("", 205206)
		return
	}
	return
}

// SetLatestProjects 设置最近的项目
func (d Doc) SetLatestProjects(
	db *gorm.DB,
	userID string,
	latestProjectID string,
) (err error) {
	var (
		count                  int64
		latestProjectIDs       string
		latestProjectIDsArr    []string
		newLatestProjectIDsArr []string
		userInfo               models.Users
	)

	// 判断最近的这个项目是否存在
	db.Table("projects").
		Where("id = ?", latestProjectID).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 204003)
		return
	}
	newLatestProjectIDsArr = append(newLatestProjectIDsArr, latestProjectID)
	// 查询历史最近项目
	db.Table("users").
		Where("id = ?", userID).
		First(&userInfo)
	latestProjectIDs = userInfo.LatestProjects

	// 移除历史记录已存在的要新增的最近项目
	latestProjectIDs = strings.Replace(latestProjectIDs, latestProjectID, "", -1)
	// 字符串转切片
	if latestProjectIDs != "" {
		latestProjectIDsArr = strings.Split(latestProjectIDs, ",")
	}

	// 最多保存3个
	for _, projectID := range latestProjectIDsArr {
		if projectID != "" && len(newLatestProjectIDsArr) < 3 {
			newLatestProjectIDsArr = append(newLatestProjectIDsArr, projectID)
		}
	}

	userInfo.LatestProjects = strings.Join(newLatestProjectIDsArr, ",")

	db.Model(&models.Users{}).
		Where("id = ?", userID).
		Updates(&userInfo)
	return
}

// SetLatestChapters 设置最近的章节
func (d Doc) SetLatestChapters(db *gorm.DB,
	latestChapterID string,
	userID string,
) (err error) {
	var (
		count                  int64
		latestChapterIDs       string
		latestChapterIDsArr    []string
		newLatestChapterIDsArr []string
		userInfo               models.Users
	)

	// 判断最近的这个章节是否存在
	db.Table("chapters").
		Where("id = ?", latestChapterID).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 204003)
		return
	}
	newLatestChapterIDsArr = append(newLatestChapterIDsArr, latestChapterID)

	// 查询历史最近章节
	db.Table("users").
		Where("id = ?", userID).
		First(&userInfo)
	latestChapterIDs = userInfo.LatestChapters

	// 移除历史记录已存在的要新增的最近章节
	latestChapterIDs = strings.Replace(latestChapterIDs, latestChapterID, "", -1)
	// 字符串转切片
	if latestChapterIDs != "" {
		latestChapterIDsArr = strings.Split(latestChapterIDs, ",")
	}

	// 最多保存6个
	for _, chapterID := range latestChapterIDsArr {
		if chapterID != "" && len(newLatestChapterIDsArr) < 6 {
			newLatestChapterIDsArr = append(newLatestChapterIDsArr, chapterID)
		}
	}

	userInfo.LatestChapters = strings.Join(newLatestChapterIDsArr, ",")

	db.Model(&models.Users{}).
		Where("id = ?", userID).
		Updates(&userInfo)
	return
}

// GetLatestProjectsAndChapters 获取最近的项目和章节
//
// @param adminID 操作人id
func (d Doc) GetLatestProjectsAndChapters(db *gorm.DB,
	adminID string,
) (projects interface{}, chapters interface{}, err error) {
	type ChapterInfo struct {
		models.Chapters
		DocName     string `json:"doc_name"`
		ProjectID   string `json:"project_id"`
		ProjectName string `json:"project_name"`
	}
	var (
		userInfo         models.Users
		projectList      []models.Projects
		chapterList      []ChapterInfo
		latestProjectIDs string
		latestChapterIDs string
	)

	// 获取记录
	db.Table("users").
		Where("id = ?", adminID).
		First(&userInfo)
	latestProjectIDs = strings.Trim(userInfo.LatestProjects, ",")
	latestChapterIDs = strings.Trim(userInfo.LatestChapters, ",")

	db.Table("projects").
		Select("id,name").
		Where("id IN (?)", strings.Split(latestProjectIDs, ",")).
		Find(&projectList)

	projects = projectList

	db.Table("chapters as c").
		Joins("left join docs d on d.id = c.doc_id").
		Joins("left join projects p on p.id = d.project_id").
		Select("c.id,c.name,c.doc_type,d.id as doc_id,d.name as doc_name,p.id as project_id,p.name as project_name").
		Where("c.id IN (?)", strings.Split(latestChapterIDs, ",")).
		Find(&chapterList)
	chapters = chapterList
	return
}

// HasAccessAuthority 验证文档权限
func (d Doc) HasAccessAuthority(db *gorm.DB, docID string, adminID string) (err error) {
	var (
		count int64
		doc   models.Docs
	)

	db.Table("docs").
		Where("id = ?", docID).
		Count(&count)
	if count <= 0 {
		return
	}
	db.Table("docs").
		Where("id = ?", docID).
		First(&doc)

	// 团队文档
	if doc.TeamID != "" && doc.TeamID != "0" {

		db.Table("teams").Where("id = ?", doc.TeamID).Count(&count)
		if count > 0 {
			// 确定团队访问权限
			t := Team{}
			_, err = t.HasAccessAuthority(db, doc.TeamID, adminID)
			if err != nil {
				// 无文档访问权限
				err = errorx.New("无文档访问权限", 205004)
			}
			return
		}
	}

	// 项目文档
	if doc.ProjectID != "" && doc.ProjectID != "0" {
		db.Table("projects").Where("id = ?", doc.ProjectID).Count(&count)
		if count > 0 {
			// 确定项目访问权限
			p := Project{}
			err = p.HasAccessAuthority(db, doc.ProjectID, adminID)
			if err != nil {
				// 无文档访问权限
				err = errorx.New("无文档访问权限", 205004)
			}
			return
		}

	}

	// 个人文档
	if doc.UserID != "" && doc.UserID != "0" {
		if doc.UserID == adminID {
			return
		}
	}
	// 无文档访问权限
	err = errorx.New("无文档访问权限", 205004)
	return
}
