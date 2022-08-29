package controllers

import (
	"znj/config"
	"znj/logics"
	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/itmisx/errorx"
	"github.com/itmisx/response"
)

// Team 团队操作
type Team struct{}

// List 团队列表
func (t *Team) List(c *gin.Context) {
	// 获取数据库连接
	db := config.NewDB()

	// 执行查询
	list, err := logics.Team{}.List(db)

	// 返回
	response.JSON(c, list, err)
}

// Save 保存（含新增和编辑）
func (t *Team) Save(c *gin.Context) {
	var team models.Teams

	// 绑定输入数据
	c.ShouldBind(&team)

	// 获取数据库连接
	db := config.NewDB()
	tx := db.Begin()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 保存数据
	err = logics.Team{}.Save(tx, team, uid)

	// 提交事务
	if err != nil {
		tx.Rollback()
		response.JSON(c, "", err)
		return
	}
	// 提交事务
	tx.Commit()
	response.JSON(c, "", err)
}

// MemberList 获取团队成员
func (t *Team) MemberList(c *gin.Context) {
	var param struct {
		ProjectID string `json:"project_id"`
		TeamID    string `json:"team_id"`
	}

	// 绑定输入参数
	_ = c.ShouldBind(&param)

	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 执行查询
	list, userRole, err := logics.Team{}.MemberList(db, param.TeamID, param.ProjectID, uid)

	response.JSON(c, gin.H{
		"user_role": userRole,
		"list":      list,
	}, err)
}

// AddTeamMember 添加团队成员
func (t *Team) AddTeamMember(c *gin.Context) {
	var user models.Users
	var param struct {
		TeamID string `json:"team_id"`
	}

	// 绑定输入参数
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.ShouldBindBodyWith(&param, binding.JSON)

	// 获取数据库连接
	db := config.NewDB()
	tx := db.Begin()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 添加逻辑
	err = logics.Team{}.AddTeamMember(tx, user, param.TeamID, uid)

	// 错误回滚
	if err != nil {
		tx.Rollback()
		response.JSON(c, "", err)
		return
	}
	// 提交，返回
	tx.Commit()
	response.JSON(c, "", err)
}

// SetWorkTeam 设置工作团队
func (t *Team) SetWorkTeam(c *gin.Context) {
	var param struct {
		WorkTeanID string `json:"work_team_id"`
	}

	// 获取数据库连接
	db := config.NewDB()

	// 绑定输入参数
	c.ShouldBind(&param)

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	// 设置逻辑
	err = logics.Team{}.SetWorkTeam(db, param.WorkTeanID, uid)

	// ajax返回
	response.JSON(c, "", err)
}

// GetWorkTeam 设置工作团队
func (t *Team) GetWorkTeam(c *gin.Context) {
	// 获取数据库连接
	db := config.NewDB()

	// 获取操作人id
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil {
		response.JSON(c, "", err)
		return
	}

	teamID := logics.Team{}.GetWorkTeam(db, uid)

	if teamID == "" {
		err = errorx.New("", 202208)
		response.JSON(c, teamID, err)
		return
	}
	// 返回
	response.JSON(c, teamID, err)
}

// ResetPassword 设为管理员
func (t *Team) ResetPassword(c *gin.Context) {
	var param struct {
		UserID   string `json:"id"`
		Password string `json:"password"`
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

	err = logics.Team{}.ResetPassword(db, param.UserID, param.Password, uid)

	response.JSON(c, "", err)
}

// SetAdminRole 设为管理员
func (t *Team) SetAdminRole(c *gin.Context) {
	var param struct {
		UserID string `json:"id"`
		TeamID string `json:"team_id"`
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

	err = logics.Team{}.SetAdminRole(db, param.UserID, param.TeamID, uid)

	response.JSON(c, "", err)
}
