package controllers

import (
	"znj/config"
	"znj/logics"
	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/response"
)

// Todo 待办操作
type Todo struct{}

// List 待办列表
func (td *Todo) List(c *gin.Context) {
	var todo models.Todos

	// 绑定输入参数
	_ = c.ShouldBind(&todo)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}
	if todo.TodoTypeID == "" {
		td.homeInfo(c)
		return
	}
	list, err := logics.Todo{}.List(db, todo.TodoTypeID, todo.DoneStatus, todo.AssignTo, uid)
	response.JSON(c, list, err)
}

// homeInfo 首页信息
func (td *Todo) homeInfo(c *gin.Context) {
	var (
		todo           models.Todos
		todoLists      interface{}
		latestProjects interface{}
		latestDocs     interface{}
		syncChan       = make(chan int, 2)
	)
	// 绑定输入参数
	_ = c.ShouldBind(&todo)
	// 获取数据库连接
	db := config.NewDB()
	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}
	// 获取待办
	go func() {
		todoLists, _ = logics.Todo{}.List(db, todo.TodoTypeID, todo.DoneStatus, todo.AssignTo, uid)
		syncChan <- 1
	}()
	// 获取最近访问的项目和文章
	go func() {
		latestProjects, latestDocs, _ = logics.Doc{}.GetLatestProjectsAndChapters(db, uid)
		syncChan <- 1
	}()
	_, _ = <-syncChan, <-syncChan
	response.JSON(c, gin.H{
		"todo_lists":      todoLists,
		"latest_projects": latestProjects,
		"latest_chapters": latestDocs,
	}, err)
}

// Save 保存待办
func (td *Todo) Save(c *gin.Context) {
	var todo models.Todos

	// 绑定输入参数
	_ = c.ShouldBind(&todo)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	err = logics.Todo{}.Save(db, todo, uid)

	response.JSON(c, "", err)
}

// GetHistory 待办历史记录
func (td *Todo) GetHistory(c *gin.Context) {
	var param struct {
		RelatedID string `json:"related_id"`
	}
	// 绑定输入参数
	_ = c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	list, err := logics.Todo{}.GetHistory(db, param.RelatedID)

	response.JSON(c, list, err)
}
