package models

// Users 用户表
type Users struct {
	ID             string `json:"id,omitempty" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	TeamID         string `json:"team_id" gorm:"column:team_id;type:varchar(30);comment:所属团队"`
	Name           string `json:"name,omitempty" gorm:"column:name;type:varchar(30);comment:昵称"`
	User           string `json:"user" validate:"required" gorm:"column:user;type:varchar(30);comment:用户名"`
	Password       string `json:"password,omitempty" validate:"required" gorm:"column:password;type:varchar(50);comment:密码"`
	CreateTime     int64  `json:"create_time,omitempty" gorm:"column:create_time;type:bigint;comment:创建时间"`
	WorkTeamID     string `json:"work_team_id" gorm:"column:work_team_id;type:varchar(30);comment:当前工作的团队"`
	LatestProjects string `json:"latest_projects" gorm:"column:latest_projects;type:varchar(30);comment:最近项目"`
	LatestChapters string `json:"latest_chapters" gorm:"column:latest_chapters;type:varchar(30);comment:最近章节"`
}

func (Users) TableName() string {
	return "users"
}

func init() {
	Models = append(Models, &Users{})
}
