package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// TodoRouter 待办路由
func TodoRouter(r *gin.RouterGroup) {
	todo := controllers.Todo{}
	// 待办列表
	r.POST("/list", func(c *gin.Context) {
		todo.List(c)
	})
	// 待办保存（含新增或编辑）
	r.POST("/save", func(c *gin.Context) {
		todo.Save(c)
	})
	// 待办变动历史
	r.POST("/history", func(c *gin.Context) {
		todo.GetHistory(c)
	})
}
