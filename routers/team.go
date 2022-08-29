package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// TeamRouter 团队路由
func TeamRouter(r *gin.RouterGroup) {
	team := controllers.Team{}
	// 团队列表
	r.POST("/list", func(c *gin.Context) {
		team.List(c)
	})
	// 团队成员
	r.POST("/memberList", func(c *gin.Context) {
		team.MemberList(c)
	})
	// 添加团队成员
	r.POST("/addTeamMember", func(c *gin.Context) {
		team.AddTeamMember(c)
	})
	// 团队保存
	r.POST("/save", func(c *gin.Context) {
		team.Save(c)
	})
	// 设置工作团队
	r.POST("/setWorkTeam", func(c *gin.Context) {
		team.SetWorkTeam(c)
	})
	// 获取工作团队
	r.POST("/getWorkTeam", func(c *gin.Context) {
		team.GetWorkTeam(c)
	})
	// 重置密码
	r.POST("/resetPassword", func(c *gin.Context) {
		team.ResetPassword(c)
	})
	// 设为管理员
	r.POST("/setAdminRole", func(c *gin.Context) {
		team.SetAdminRole(c)
	})
}
