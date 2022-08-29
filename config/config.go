package config

import (
	"github.com/itmisx/gormx"
	"github.com/itmisx/logger"
	"github.com/itmisx/redisx"
	"github.com/spf13/viper"
)

var Conf struct {
	HTTPPort     string        `mapstructure:"http_port"`
	Mysql        gormx.Config  `mapstructure:"mysql"`
	Redis        redisx.Config `mapstructure:"redis"`
	Logger       logger.Config `mapstructure:"logger"`
	DefalutLang  string        `mapstructure:"default_lang"`
	AllowMethods []string      `mapstructure:"allow_methods"`
	AllowHeaders []string      `mapstructure:"allow_headers"`
	AdminUser    struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"admin_user"`
}

// 解析配置
func ParseConf() {
	vp := viper.New()
	vp.SetConfigName("znj")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./etc")
	vp.AddConfigPath("/etc")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	vp.Unmarshal(&Conf)
}
