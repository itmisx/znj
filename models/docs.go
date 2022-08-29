package models

// Docs 文档类型
type Docs struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	Name       string `json:"name" gorm:"column:name;type:varchar(50);comment:文档标题"`
	TeamID     string `json:"team_id" gorm:"column:team_id;type:varchar(30);comment:所属团队"`
	ProjectID  string `json:"project_id" gorm:"column:project_id;type:varchar(30);comment:所属项目"`
	UserID     string `json:"user_id" gorm:"column:user_id;type:varchar(30);comment:创建人"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`
}

func (Docs) TableName() string {
	return "docs"
}

func init() {
	Models = append(Models, &Docs{})
}
