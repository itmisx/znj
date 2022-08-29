package migrate

import (
	"errors"
	"time"

	"znj/config"
	"znj/define"
	"znj/models"

	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// Install 数据表安装
func Install() error {
	var err error
	db := config.NewDB()
	db.Transaction(func(tx *gorm.DB) error {
		var transErr error
		// 初始化表结构
		transErr = tx.Set("gorm:table_options", "ENGINE=InnoDB").
			Set("gorm:table_options", "CHARSET=utf8mb4").
			Set("gorm:table_options", "COLLATE=utf8mb4_bin").
			AutoMigrate(models.Models...)
		if transErr != nil {
			err = transErr
			return err
		}
		// 创建默认的管理员及团队
		{
			// 创建管理员账号
			admin := models.Users{}
			admin.ID, transErr = (helper.UUID{}).String()
			if transErr != nil {
				err = transErr
				return err
			}
			admin.Name = config.Conf.AdminUser.Username
			admin.User = config.Conf.AdminUser.Username
			admin.Password = helper.Md5(config.Conf.AdminUser.Password)
			admin.CreateTime = time.Now().Unix()
			rowsAffected := db.Model(&models.Users{}).
				Create(&admin).
				RowsAffected
			if rowsAffected <= 0 {
				err = errors.New("create admin user failed")
				return err
			}
			// 创建默认团队
			var team models.Teams
			if ID, err := (helper.UUID{}).String(); err != nil {
				err = errors.New("create default team failed")
				return err
			} else {
				team.ID = ID
			}
			team.Name = "我的团队"
			team.CreateID = admin.ID
			team.CreateTime = time.Now().Unix()
			rowsAffected = db.Model(&models.Teams{}).Create(&team).RowsAffected
			if rowsAffected <= 0 {
				err = errors.New("create admin user failed")
				return err
			}
			// 团队成员
			var teamMember models.TeamMember
			if ID, err := (helper.UUID{}).String(); err != nil {
				err = errors.New("create teamMember failed")
				return err
			} else {
				teamMember.ID = ID
			}
			teamMember.TeamID = team.ID
			teamMember.UserID = admin.ID
			teamMember.Role = 1
			teamMember.CreateTime = time.Now().Unix()
			rowsAffected = db.Model(&models.TeamMember{}).Create(&teamMember).RowsAffected
			if rowsAffected <= 0 {
				err = errors.New("create teamMember failed")
				return err
			}
		}
		// 插入数据库版本信息
		{
			var version models.Version
			rs := tx.Model(&models.Version{}).Attrs(models.Version{
				Version: define.DBVersion,
			}).FirstOrCreate(&version)
			if rs.Error != nil {
				return rs.Error
			}
		}
		return nil
	})
	return err
}
