package models

// 数据库版本
type Version struct {
	Version string `json:"version" gorm:"column:version;primaryKey;type:varchar(50) NOT NULL;comment:版本号"`
}

func (Version) TableName() string {
	return "version"
}

func init() {
	Models = append(Models, &Version{})
}
