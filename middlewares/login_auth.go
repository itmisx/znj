package middlewares

import (
	"znj/pkg"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/errorx"
	"github.com/itmisx/response"
)

// IsLogin 登录判断
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否已经登录
		uid, err := pkg.SessionGet(c, "uid")
		if err != nil || uid == "" {
			err = errorx.New("", 200000)
			response.JSON(c, "", err)
			c.Abort()
			return
		}
		c.Next()
	}
}
