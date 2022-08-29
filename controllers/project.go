package controllers

import (
	"znj/config"
	"znj/logics"
	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/response"
)

// Project 项目操作
type Project struct{}

// List 项目列表
func (pj *Project) List(c *gin.Context) {
	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 获取工作团队
	workTeamID := logics.Team{}.GetWorkTeam(db, uid)

	// 执行查询
	list, err := logics.Project{}.List(db, workTeamID, uid)

	response.JSON(c, list, err)
}

// Save 项目编辑 (含新增和编辑)
func (pj *Project) Save(c *gin.Context) {
	var project models.Projects

	// 绑定输入参数
	c.ShouldBind(&project)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 获取工作团队
	workTeamID := logics.Team{}.GetWorkTeam(db, uid)

	// 保存数据
	err = logics.Project{}.Save(db, workTeamID, project, uid)

	response.JSON(c, "", err)
}

// TodoTypeSave 待办类型保存（包含新增和编辑）
func (pj *Project) TodoTypeSave(c *gin.Context) {
	var todoType models.TodoType

	// 绑定输入参数
	c.ShouldBind(&todoType)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 保存数据
	err = logics.Project{}.TodoTypeSave(db, todoType, uid)

	response.JSON(c, "", err)
}

// TodoTypeList 待办类型列表
func (pj *Project) TodoTypeList(c *gin.Context) {
	var param struct {
		ProjectID string `json:"project_id"`
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
	list, err := logics.Project{}.TodoTypeList(db, param.ProjectID, uid)

	response.JSON(c, list, err)
}

// MemberList 项目成员列表
func (pj *Project) MemberList(c *gin.Context) {
	var param struct {
		ProjectID string `json:"project_id"`
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
	list, err := logics.Project{}.MemberList(db, param.ProjectID, uid)

	response.JSON(c, list, err)
}

// MemberSet 项目成员设置
func (pj *Project) MemberSet(c *gin.Context) {
	var param struct {
		ProjectID string `json:"project_id"`
		IDs       string `json:"ids"`
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

	// 保存数据
	err = logics.Project{}.MemberSet(db, param.ProjectID, param.IDs, uid)

	response.JSON(c, "", err)
}

// TodoTypeBatchSort 待办类型批量排序
func (pj *Project) TodoTypeBatchSort(c *gin.Context) {
	var param struct {
		IDSorts string `json:"id_sorts"`
	}
	// 绑定输入参数
	c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 保存数据
	err := logics.Project{}.TodoTypeBatchSort(db, param.IDSorts)

	response.JSON(c, "", err)
}
