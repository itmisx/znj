package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// IndexRouter index路由
func IndexRouter(r *gin.RouterGroup) {
	// 是否登录
	r.POST("/isSignedIn", func(c *gin.Context) {
		controllers.IsSignedIn(c)
	})
	// 登录
	r.POST("/signin", func(c *gin.Context) {
		controllers.Signin(c)
	})
	// 登出
	r.POST("/signout", func(c *gin.Context) {
		controllers.Signout(c)
	})
}
