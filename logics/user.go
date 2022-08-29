package logics

import (
	"regexp"
	"strings"

	"znj/models"
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// User 用户操作
type User struct{}

// GetUserID 获取当前登录的用户id（从session中获取）
func (u User) GetUserID(c *gin.Context) (userID string, err error) {
	uid, err := pkg.SessionGet(c, "uid")
	// 判断当前用户是否登录，未登录的直接提示登录
	if err != nil || uid == "" {
		err = errorx.New("", 200000)
		return
	}
	return
}

// CreateUser 创建用户
func (u User) CreateUser(db *gorm.DB,
	user models.Users,
	teamID string,
) (userID string, err error) {
	var (
		primaryKeyID string = ""
		count        int64  = 0
		rowsAffected int64  = 0
	)
	// 主键id
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	user.ID = primaryKeyID
	// 用户名
	if user.Name == "" {
		err = errorx.New("", 202201)
		return // 姓名不能为空
	}
	// 登录账号
	if user.User == "" {
		err = errorx.New("", 202202)
		return // 账号不能为空
	}
	// 密码
	if user.Password == "" {
		err = errorx.New("", 202203)
		return // 密码不能为空
	}

	user.TeamID = teamID

	// 判断密码是否包含大小写字母和数字
	{
		regexpRes1, _ := regexp.MatchString("[A-Z]+", user.Password)  // 大写字母
		regexpRes2, _ := regexp.MatchString("[a-z]+", user.Password)  // 小写字母
		regexpRes3, _ := regexp.MatchString("[0-9]+", user.Password)  // 数字
		regexpRes4, _ := regexp.MatchString("(.){8,}", user.Password) // 长度不能低于8位
		if !regexpRes1 || !regexpRes2 || !regexpRes3 || !regexpRes4 {
			err = errorx.New("", 201006)
			return
		}
	}
	user.Password = helper.Md5(helper.Md5(user.Password))
	// 判断用户是否已经存在
	db.Table("users").
		Where("user = ?", user.User).
		Count(&count)
	if count > 0 {
		err = errorx.New("", 202204)
		return // 用户已存在
	}

	// 添加用户
	rowsAffected = db.Create(&user).RowsAffected

	if rowsAffected <= 0 {
		err = errorx.New("", 202205)
		return // 添加用户失败
	}
	userID = primaryKeyID
	return
}

// ChangePassword 修改密码
//
// @param adminID 操作人id
// @param oldPassword 旧密码
// @param newPassword 新密码
// @param newPasswordConfirm 确认密码
func (u User) ChangePassword(
	db *gorm.DB,
	oldPassword string,
	newPassword string,
	newPasswordConfirm string,
	adminID string,
) (err error) {
	var (
		count int64
		user  models.Users
	)

	/** 字段验证
	 */
	if strings.TrimSpace(oldPassword) == "" {
		err = errorx.New("", 201001)
		return
	}
	if strings.TrimSpace(newPassword) == "" {
		err = errorx.New("", 201002)
		return
	}
	if strings.TrimSpace(newPasswordConfirm) == "" {
		err = errorx.New("", 201003)
		return
	}
	if strings.TrimSpace(newPassword) != strings.TrimSpace(newPasswordConfirm) {
		err = errorx.New("", 201005)
		return
	}

	// 判断密码是否包含大小写字母和数字
	{
		regexpRes1, _ := regexp.MatchString("[A-Z]+", newPassword)  // 大写字母
		regexpRes2, _ := regexp.MatchString("[a-z]+", newPassword)  // 小写字母
		regexpRes3, _ := regexp.MatchString("[0-9]+", newPassword)  // 数字
		regexpRes4, _ := regexp.MatchString("(.){8,}", newPassword) // 长度不能低于8位
		if regexpRes1 && regexpRes2 && regexpRes3 && !regexpRes4 {
			err = errorx.New("", 201006)
			return
		}
	}
	oldPassword = helper.Md5(helper.Md5(oldPassword))
	// 判断原始密码是否正确
	db.Model(user).Where("id = ? and password = ?", adminID, oldPassword).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 201004)
		return
	}
	newPassword = helper.Md5(helper.Md5(newPassword))
	// 保存至数据库
	{
		// 更新密码
		rowsAffected := db.Model(user).
			Where("id = ?", adminID).
			Update("password", newPassword).
			RowsAffected
		if rowsAffected != 1 {
			err = errorx.New("", 201007)
			return
		}
	}

	// 返回错误码
	return
}
