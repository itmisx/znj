package logics

import (
	"time"

	"znj/config"
	"znj/models"
	"znj/pkg"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
)

// LoginUserInfo 登录信息
type LoginUserInfo struct {
	models.Users
	Token string `json:"token"`
}

// IsSignedIn 判断是否已经登录
func IsSignedIn(c *gin.Context) (err error) {
	uid, err := pkg.SessionGet(c, "uid")
	if err != nil || uid == "" {
		err = errorx.New("", 200000)
		return
	}
	var count int64
	config.NewDB().Model(&models.Users{}).Where("id = ?", uid).Count(&count)
	if count <= 0 {
		err = errorx.New("", 200000)
		return
	}
	return
}

// Signin 登录
func Signin(db *gorm.DB,
	user models.Users,
) (userInfo LoginUserInfo, err error) {
	var count int64 = 0

	// 字段验证
	validate := validator.New()
	err1 := validate.Struct(&user)
	if err1 != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "User" && e.ActualTag() == "required" {
				err = errorx.New("", 200002)
				return
			} else if e.Field() == "Password" && e.ActualTag() == "required" {
				err = errorx.New("", 200003)
				return
			} else {
				err = errorx.New("", 200001)
				return
			}
		}
	}
	fields := "id,name,user,create_time"
	user.Password = helper.Md5(user.Password)

	db.Table("users").Select(fields).
		Where("user = ? and password = ?", user.User, user.Password).
		Find(&userInfo).
		Count(&count)
	if count != 1 {
		err = errorx.New("", 200001)
		return
	}
	// 生成token
	userInfo.Token, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	userInfo.Token = helper.Md5(userInfo.Token)
	return
}

// SaveHistory 保存操作历史
// historyType 记录类型
// operationType 操作类型
// relatedID 关联资源
// userID 操作人
func SaveHistory(db *gorm.DB,
	historyType string,
	operationType string,
	relatedID string,
	userID string,
) (errCode int64) {
	var (
		rowsAffected int64
		primaryKeyID string
		History      models.History
	)
	// 获取唯一主键
	primaryKeyID, err := helper.UUID{}.String()
	if err != nil {
		return
	}
	History.ID = primaryKeyID
	History.HistoryType = historyType
	History.OperationType = operationType
	History.RelatedID = relatedID
	History.CreateID = userID
	History.CreateTime = time.Now().Unix()

	rowsAffected = db.Table("history").Create(History).RowsAffected
	if rowsAffected != 1 {
		errCode = 100003
		return
	}
	return
}
