package models

// Teams 团队表
type Teams struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	Name       string `json:"name" gorm:"column:name;type:varchar(50);comment:团队名称"`
	Desc       string `json:"desc" gorm:"column:desc;type:varchar(200);comment:团队描述"`
	CreateID   string `json:"create_id" gorm:"column:create_id;type:varchar(30);comment:成员id"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`
}

func (Teams) TableName() string {
	return "teams"
}

func init() {
	Models = append(Models, &Teams{})
}
