package models

// Projects 项目表
type Projects struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	Name       string `json:"name" gorm:"column:name;type:varchar(50);comment:项目名称"`
	TeamID     string `json:"team_id" gorm:"column:team_id;type:varchar(50);comment:所属团队"`
	Desc       string `json:"desc" gorm:"column:desc;type:varchar(50);comment:项目描述"`
	CreateID   string `json:"create_id" gorm:"column:create_id;type:varchar(30);comment:创建人"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`
	StarTime   int64  `json:"star_time" gorm:"column:star_time;type:bigint;comment:置顶时间"`
}

func (Projects) TableName() string {
	return "projects"
}

func init() {
	Models = append(Models, &Projects{})
}
