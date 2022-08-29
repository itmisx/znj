package main

import (
	// 通用操作方法
	// 路由初始化

	"embed"
	"io/fs"
	"log"
	"net/http"

	"znj/config"
	"znj/define"
	"znj/define/lang"
	"znj/migrate"

	"znj/middlewares"

	"znj/routers"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/go-helper"
	"github.com/itmisx/i18n"
	"github.com/itmisx/logger"
)

//go:embed web_src
var web embed.FS

func main() {
	// config配置初始化
	config.ParseConf()
	// 加载语言包
	i18n.LoadLangPack(lang.ErrCode)
	// logger
	logger.Init(config.Conf.Logger, logger.String("service.name", "znj"), logger.String("service.version", define.Version))
	// redis初始化
	config.RedisInit(config.Conf.Redis)
	// mysql初始化
	config.MysqlInit(config.Conf.Mysql)
	// uuid 初始化
	helper.UUID{}.Init(config.Conf.Redis)
	// 数据库迁移
	migrate.AutoMigrate()
	r := gin.Default()
	// 允许跨域
	r.Use(middlewares.Cors())
	r.Use(middlewares.Download())
	// 加载web资源
	embedSub, _ := fs.Sub(web, "web_src/dist")
	r.StaticFS("/", http.FS(embedSub))
	// 路由分组
	routers.RouterGroup(r)
	// 在指定端口启动http服务
	err := r.Run(config.Conf.HTTPPort)
	log.Println(err)
}
