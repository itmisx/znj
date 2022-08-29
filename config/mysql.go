package config

import (
	"github.com/itmisx/gormx"
	"gorm.io/gorm"
)

var _db *gorm.DB

// MysqlInit mysql数据库初始化
func MysqlInit(cfg gormx.Config) (err error) {
	_db, err = gormx.New(cfg)
	if err != nil {
		panic(err)
	}
	return
}

// 获取一个数据库对象
func NewDB() *gorm.DB {
	return _db
}
