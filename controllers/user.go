package controllers

import (
	"fmt"

	"znj/config"
	"znj/logics"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/response"
)

// User 控制器
type User struct{}

// ChangePassword 修改密码
func (u *User) ChangePassword(c *gin.Context) {
	var param struct {
		OldPassword        string `json:"old_password"`
		NewPassword        string `json:"new_password"`
		NewPasswordConfirm string `json:"new_password_confirm"`
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
	fmt.Printf("%+v", param)
	err = logics.User{}.ChangePassword(db, param.OldPassword,
		param.NewPassword,
		param.NewPasswordConfirm,
		uid)

	response.JSON(c, "", err)
}
