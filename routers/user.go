package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
func UserRouter(r *gin.RouterGroup) {
	user := controllers.User{}
	// 修改密码
	r.POST("/changePassword", func(c *gin.Context) {
		user.ChangePassword(c)
	})
}
