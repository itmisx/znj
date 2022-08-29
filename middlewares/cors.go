package middlewares

import (
	"znj/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOriginFunc = func(origin string) bool {
		// TODO config
		return true
	}
	conf.AddAllowMethods(config.Conf.AllowMethods...)
	conf.AddAllowHeaders(config.Conf.AllowHeaders...)
	conf.AllowCredentials = true
	return cors.New(conf)
}
