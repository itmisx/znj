package routers

import (
	"znj/middlewares"

	"github.com/gin-gonic/gin"
)

// RouterGroup 自定义路由分组
func RouterGroup(r *gin.Engine) {
	// index路由
	index := r.Group("/index")
	IndexRouter(index)

	// 登录鉴权中间件
	r.Use(middlewares.IsLogin())

	// index路由
	user := r.Group("/user")
	UserRouter(user)

	// todo路由
	todo := r.Group("/todo")
	TodoRouter(todo)

	// project路由
	project := r.Group("/project")
	ProjectRouter(project)

	// team路由
	doc := r.Group("/doc")
	DocRouter(doc)

	// team路由
	team := r.Group("/team")
	TeamRouter(team)

	// 文件上传
	upload := r.Group("/upload")
	UploadRouter(upload)
}
