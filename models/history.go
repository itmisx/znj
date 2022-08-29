package models

// History 历史操作
type History struct {
	ID            string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	RelatedID     string `json:"related_id" gorm:"column:related_id;type:varchar(30);comment:关联资源id"`
	HistoryType   string `json:"history_type"  gorm:"column:history_type;type:varchar(30);comment:记录类型"`
	OperationType string `json:"operation_type" gorm:"column:operation_type;type:varchar(30);comment:操作类型"`
	CreateID      string `json:"create_id" gorm:"column:create_id;type:varchar(30);comment:操作人"`
	CreateTime    int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
}

func (History) TableName() string {
	return "history"
}

func init() {
	Models = append(Models, &History{})
}
