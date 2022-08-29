package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// ProjectRouter 项目路由
func ProjectRouter(r *gin.RouterGroup) {
	project := controllers.Project{}
	// 项目列表
	r.POST("/list", func(c *gin.Context) {
		project.List(c)
	})
	// 项目添加或编辑保存
	r.POST("/save", func(c *gin.Context) {
		project.Save(c)
	})
	//待办类型保存
	r.POST("/todoTypeSave", func(c *gin.Context) {
		project.TodoTypeSave(c)
	})
	// 待办类型列表
	r.POST("/todoTypeList", func(c *gin.Context) {
		project.TodoTypeList(c)
	})
	// 项目成员列表
	r.POST("/memberList", func(c *gin.Context) {
		project.MemberList(c)
	})
	// 项目成员设置
	r.POST("/memberSet", func(c *gin.Context) {
		project.MemberSet(c)
	})
	// 任务类型批量排序
	r.POST("/todoTypeBatchSort", func(c *gin.Context) {
		project.TodoTypeBatchSort(c)
	})
}
