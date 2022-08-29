package models

// TeamMember 团队成员表
type TeamMember struct {
	ID         string `json:"id" gorm:"column:id;primaryKey;type:varchar(30);comment:主键"`
	TeamID     string `json:"team_id" gorm:"column:team_id;type:varchar(30);comment:团队id"`
	UserID     string `json:"user_id" gorm:"column:user_id;type:varchar(30);comment:成员id"`
	Role       int64  `json:"role" gorm:"column:role;type:int;comment:角色,1-超级管理员，2-管理员，3-普通用户"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`
}

func (TeamMember) TableName() string {
	return "team_member"
}

func init() {
	Models = append(Models, &TeamMember{})
}
