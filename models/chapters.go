package models

// Chapters 章节模型
type Chapters struct {
	ID         string `json:"id"  gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	PID        string `json:"pid" gorm:"column:pid;type:varchar(30);comment:上级章节id"`
	DocType    int64  `json:"doc_type" gorm:"column:doc_type;type:int;default:0;comment:文档类型"`
	Sort       int64  `json:"sort" gorm:"column:sort;type:int;default:0;comment:排序"`
	Name       string `json:"name" gorm:"column:name;type:varchar(50);comment:章节标题"`
	Version    string `json:"version" gorm:"column:version;type:varchar(30);comment:最新内容版本"`
	DocID      string `json:"doc_id" gorm:"column:doc_id;type:varchar(30);comment:所属文档id"`
	CreateID   string `json:"create_id" gorm:"column:create_id;type:varchar(30);comment:创建人"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`
}

func (Chapters) TableName() string {
	return "chapters"
}

func init() {
	Models = append(Models, &Chapters{})
}
