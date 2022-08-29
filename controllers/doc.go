package controllers

import (
	"znj/config"
	"znj/logics"
	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/response"
)

// Doc 文档操作
type Doc struct{}

// List 文档列表
func (d Doc) List(c *gin.Context) {
	var (
		workTeamID string
		param      struct {
			ProjectID string `json:"project_id"`
			TeamID    string `json:"team_id"`
		}
	)

	// 获取输入变量
	c.ShouldBind(&param)

	// 获取数据库资源
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}
	// 获取当前的工作团队
	// teamID 指示是否要获取团队知识库
	if param.TeamID != "" {
		workTeamID = logics.Team{}.GetWorkTeam(db, uid)
	}

	// 执行查询
	list, err := logics.Doc{}.List(db, uid, workTeamID, param.ProjectID)

	// 返回结果
	response.JSON(c, list, err)
}

// Save 文档编辑
func (d Doc) Save(c *gin.Context) {
	var (
		doc models.Docs
		err error
	)
	// 获取输入参数
	c.ShouldBind(&doc)

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 获取数据库连接
	db := config.NewDB()

	// 获取当前的工作团队
	workTeamID := logics.Team{}.GetWorkTeam(db, uid)

	// 保存文档
	err = logics.Doc{}.Save(db, doc, uid, workTeamID)

	response.JSON(c, "", err)
}

// ChapterBatchSort 待办类型批量排序
func (d Doc) ChapterBatchSort(c *gin.Context) {
	var param struct {
		IDSorts string `json:"id_sorts"`
	}
	// 绑定输入参数
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 保存数据
	err := logics.Doc{}.ChapterBatchSort(db, param.IDSorts)

	response.JSON(c, "", err)
}

// ChapterList 章节列表
func (d Doc) ChapterList(c *gin.Context) {
	var param struct {
		DocID string `json:"doc_id"`
	}

	// 绑定输入参数
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 执行查询
	list, err := logics.Doc{}.ChapterList(db, param.DocID, uid)

	response.JSON(c, list, err)
}

// ChapterSave 章节保存
func (d Doc) ChapterSave(c *gin.Context) {
	var chapter models.Chapters

	// 输入参数绑定
	c.ShouldBind(&chapter)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 保存逻辑
	err = logics.Doc{}.ChapterSave(db, chapter, uid)

	response.JSON(c, "", err)
}

// ChapterContent 章节内容
func (d Doc) ChapterContent(c *gin.Context) {
	var chapterContent models.ChapterContent
	// 绑定输入参数
	c.ShouldBind(&chapterContent)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 执行查询
	content, err := logics.Doc{}.ChapterContent(db, chapterContent, uid)

	// 返回结果
	response.JSON(c, content, err)
}

// ChapterContentList 章节内容列表
func (d Doc) ChapterContentList(c *gin.Context) {
	var param struct {
		ChapterID string `json:"chapter_id"`
	}
	// 绑定参数
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 执行查询
	list, err := logics.Doc{}.ChapterContentList(db, param.ChapterID, uid)

	// 返回结果
	response.JSON(c, list, err)
}

// ChapterContentSave 章节内容保存
func (d Doc) ChapterContentSave(c *gin.Context) {
	var chapterContent models.ChapterContent

	// 获取数据库连接
	db := config.NewDB()
	tx := db.Begin()

	// 绑定输入内容
	c.ShouldBind(&chapterContent)

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 保存逻辑
	err = logics.Doc{}.ChapterContentSave(tx, chapterContent, uid)

	if err != nil {
		tx.Rollback()
		response.JSON(c, "", err)
		return
	}
	tx.Commit()
	response.JSON(c, "", err)
}

// GetChapterRevision 获取章节版本内容
func (d Doc) GetChapterRevision(c *gin.Context) {
	var param struct {
		ChapterID      string `json:"chapter_id"`
		ChapterVersion string `json:"version"`
	}
	// 绑定输入参数
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 获取内容
	content, err := logics.Doc{}.GetChapterRevision(db, param.ChapterID, param.ChapterVersion, uid)

	// 返回
	response.JSON(c, content, err)
}

// SetChapterRevision 章节内容回滚
func (d Doc) SetChapterRevision(c *gin.Context) {
	var param struct {
		ChapterID      string `json:"id"`
		CHapterVersion string `json:"version"`
	}

	// 绑定输入数据
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 回滚逻辑
	err = logics.Doc{}.SetChapterRevision(db, param.ChapterID, param.CHapterVersion, uid)

	// 返回结果
	response.JSON(c, "", err)
}

// GetLatestProjectsAndChapters 获取最近访问的项目和文档
func (d Doc) GetLatestProjectsAndChapters(c *gin.Context) {
	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 执行查询
	projects, chapters, err := logics.Doc{}.GetLatestProjectsAndChapters(db, uid)

	response.JSON(c, gin.H{
		"projects": projects,
		"chapters": chapters,
	}, err)
}
