package models

// ProjectMember 模型
type ProjectMember struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	ProjectID  string `json:"project_id" gorm:"column:project_id;type:varchar(30);comment:项目id"`
	UserID     string `json:"user_id" gorm:"column:user_id;type:varchar(30);comment:成员id"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}

func init() {
	Models = append(Models, &ProjectMember{})
}
