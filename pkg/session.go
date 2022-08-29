package pkg

import (
	"context"
	"time"

	"znj/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/itmisx/errorx"
)

// SessionSet 设置会话值
func SessionSet(token string, key string, value string) (err error) {
	// 实例化redis
	rc := config.NewRDB()
	// 判断sessionKey是否存在
	sessionKey := "znj:login_token:" + token
	// 如果sessionKey存在
	{
		_, err = rc.HSet(context.Background(), sessionKey, key, value).Result()
		if err != nil {
			err = errorx.New(err.Error(), 100000)
			return
		}
	}
	// 延长key的有效期
	rc.Expire(context.Background(), sessionKey, 36000*time.Second)
	return
}

// SessionClear 清除会话内容
func SessionClear(c *gin.Context) (err error) {
	var param struct {
		Authorization string `json:"Authorization"`
	}
	c.ShouldBindHeader(&param)

	rc := config.NewRDB()

	key := "znj:login_token:" + param.Authorization
	// 判断是否存在
	{
		var num int64
		num, err = rc.Del(context.Background(), key).Result()
		switch {
		case err == redis.Nil:
			err = nil
		case err != nil:
			err = errorx.New(err.Error(), 100000)
		case num == 0:
			err = errorx.New(err.Error(), 100000)
		}
	}
	return
}

// SessionGet 获取会话值
func SessionGet(c *gin.Context, key string) (value string, err error) {
	var param struct {
		Authorization string `json:"Authorization"`
	}
	c.ShouldBindHeader(&param)
	// 实例化redis
	rc := config.NewRDB()

	// 判断sessionKey是否存在
	sessionKey := "znj:login_token:" + param.Authorization
	{
		_, err := rc.Exists(context.Background(), sessionKey).Result()
		if err != nil {
			return "", errorx.New(err.Error(), 100000)
		}
	}
	// 如果sessionKey存在
	{
		rs, err := rc.HGet(context.Background(), sessionKey, key).Result()
		if err != nil {
			return "", errorx.New(err.Error(), 100000)
		}
		rc.Expire(context.Background(), sessionKey, 36000*time.Second)
		return rs, nil
	}
}
