package models

// TodoType 待办分组类型
type TodoType struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	ProjectID  string `json:"project_id" gorm:"column:project_id;type:varchar(30);comment:项目id"`
	Name       string `json:"name" gorm:"column:name;type:varchar(30);comment:项目id"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`
	StarTime   int64  `json:"star_time" gorm:"column:star_time;type:bigint;comment:置顶时间"`
	SortInit   int64  `json:"sort_init" gorm:"column:sort_init;type:bigint;comment:排序初始值"`
	Sort       int64  `json:"sort" gorm:"column:sort;type:bigint;comment:排序"`
}

func (TodoType) TableName() string {
	return "todo_type"
}

func init() {
	Models = append(Models, &TodoType{})
}
