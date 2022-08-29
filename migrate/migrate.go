package migrate

import (
	"znj/config"
	"znj/models"
)

// 升级版本结构体定义
// version+script
type versionScript struct {
	DstVersion string       // 要升级的最新版
	Func       func() error // 升级脚本
}

// 版本记录
// 记录了每个版本可升级的版本及升级脚本
var versions = map[string]versionScript{
	// key-当前版本
	// value-下个版本及对应的升级脚本
	// "1": {
	// 	DstVersion: "2", // 空代表最新版
	// 	Func:       Migrate_2,
	// },
}

// AutoMigrate 检查或更新数据库
func AutoMigrate() {
	db := config.NewDB()
	hasDB := db.Migrator().HasTable("version")
	// 首次自动安装
	if !hasDB {
		if err := Install(); err != nil {
			panic(err.Error())
		}
	} else {
		// 数据库迁移
		migrate()
	}
}

// migrate 数据迁移
// 获取当前版本，然后根据版本记录，循环升级到最新版本(直到next_version为空)
func migrate() {
	var (
		currentVer models.Version
		newVer     models.Version
	)
	// 获取当前版本
	db := config.NewDB()
	db.Model(&models.Version{}).Take(&currentVer)

	v := currentVer.Version
	// 循环升级，直到升级到最新版本
	for {
		// 获取对应版本的升级脚本
		script, ok := versions[v]
		if ok && script.DstVersion != "" {
			// 执行升级版本，如果遇到错误，则终止
			err := script.Func()
			if err != nil {
				panic(err.Error())
			}
			// 执行成功后，保存当前版本记录到数据库
			newVer.Version = script.DstVersion
			db.Model(&models.Version{}).Where("version = ?", v).Updates(&newVer)
			v = script.DstVersion
		} else {
			break
		}
	}
}
