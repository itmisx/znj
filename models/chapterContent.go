package models

// ChapterContent 章节内容
type ChapterContent struct {
	Version    string `json:"version" gorm:"column:version;primaryKey;type:varchar(30);comment:版本-主键"`
	ChapterID  string `json:"chapter_id" gorm:"column:chapter_id;type:varchar(30);comment:章节id"`
	Content    string `json:"content" gorm:"column:content;type:text;comment:章节内容"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint(20);comment:创建时间"`
	CreateID   string `json:"create_id" gorm:"column:create_id;type:varchar(30);comment:创建人"`
	Comment    string `json:"comment" gorm:"column:comment;type:varchar(200);comment:注释"`
}

func (ChapterContent) TableName() string {
	return "chapter_content"
}

func init() {
	Models = append(Models, &ChapterContent{})
}
