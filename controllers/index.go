package controllers

import (
	"znj/config"
	"znj/logics"
	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/response"
)

// IsSignedIn 是否已登录
func IsSignedIn(c *gin.Context) {
	err := logics.IsSignedIn(c)
	response.JSON(c, "", err)
}

// Signin 登录
func Signin(c *gin.Context) {
	var user models.Users
	// 绑定输入参数
	c.ShouldBind(&user)

	// 获取数据库连接
	db := config.NewDB()

	// 登录验证
	userInfo, err := logics.Signin(db, user)
	/*登录成功*/
	if err == nil {
		// 保存session
		pkg.SessionSet(userInfo.Token, "uid", userInfo.ID)
		// 返回登录结果
		response.JSON(c, userInfo, err)
	} else {
		/*登录失败*/
		response.JSON(c, "", err)
	}
}

// Signout 登出
func Signout(c *gin.Context) {
	err := pkg.SessionClear(c)
	response.JSON(c, "", err)
}

// SetLang 设置语言
func SetLang(c *gin.Context) {
	var param struct {
		Token string `json:"token"`
		Lang  string `json:"lang"`
	}
	c.ShouldBindHeader(&param)
	c.ShouldBind(&param)
	err := pkg.SessionSet(param.Token, "lang", param.Lang)
	response.JSON(c, "", err)
}
