package models

// Todos 待办表
type Todos struct {
	ID         string `json:"id" validate:"required" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	Level      int16  `json:"level" validate:"required,min=1,max=4" gorm:"column:level;type:int;comment:待办优先级"`
	Content    string `json:"content" validate:"required" gorm:"column:content;type:varchar(200);comment:待办内容"`
	Desc       string `json:"desc" gorm:"column:desc;type:text;comment:任务描述"`
	DoneStatus int16  `json:"done_status" validate:"min=1,max=4" gorm:"column:done_status;type:int;comment:完成状态"`
	TodoTypeID string `json:"todo_type_id" gorm:"column:todo_type_id;type:varchar(30);comment:任务id"`
	ProjectID  string `json:"project_id" gorm:"column:project_id;type:varchar(30);comment:项目id"`
	AssignTo   string `json:"assign_to" validate:"required" gorm:"column:assign_to;type:varchar(30);comment:分派"`
	CreateID   string `json:"create_id" validate:"required" gorm:"column:create_id;type:varchar(30);comment:创建人"`
	CreateTime int64  `json:"create_time" validate:"required" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" validate:"required" gorm:"column:update_time;type:bigint;comment:更新时间"`
	EndTime    string `json:"end_time" gorm:"column:end_time;type:bigint;comment:更新时间"`
}

func (Todos) TableName() string {
	return "todos"
}

func init() {
	Models = append(Models, &Todos{})
}
