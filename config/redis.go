package config

import "github.com/itmisx/redisx"

// 实例变量
var _cli redisx.Client

// RedisInit
func RedisInit(conf redisx.Config) {
	_cli = redisx.New(conf)
}

// NewRDB 获取redis链接
func NewRDB() redisx.Client {
	return _cli
}
