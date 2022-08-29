package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// DocRouter 文档路由
func DocRouter(r *gin.RouterGroup) {
	doc := controllers.Doc{}
	// 文档列表
	r.POST("/list", func(c *gin.Context) {
		doc.List(c)
	})
	// 文档保存
	r.POST("/save", func(c *gin.Context) {
		doc.Save(c)
	})
	// 章节批量排序
	r.POST("/chapterBatchSort", func(c *gin.Context) {
		doc.ChapterBatchSort(c)
	})
	// 文档章节列表
	r.POST("/chapterList", func(c *gin.Context) {
		doc.ChapterList(c)
	})
	// 文档章节保存
	r.POST("/chapterSave", func(c *gin.Context) {
		doc.ChapterSave(c)
	})
	// 文档内容
	r.POST("/chapterContent", func(c *gin.Context) {
		doc.ChapterContent(c)
	})
	// 文档内容列表
	r.POST("/chapterContentList", func(c *gin.Context) {
		doc.ChapterContentList(c)
	})
	// 文档内容保存
	r.POST("/chapterContentSave", func(c *gin.Context) {
		doc.ChapterContentSave(c)
	})
	// 文档内容保存
	r.POST("/getChapterRevision", func(c *gin.Context) {
		doc.GetChapterRevision(c)
	})
	// 文档内容保存
	r.POST("/setChapterRevision", func(c *gin.Context) {
		doc.SetChapterRevision(c)
	})
	// GetLatestProjectsAndChapters
	r.POST("/getLatestProjectsAndChapters", func(c *gin.Context) {
		doc.GetLatestProjectsAndChapters(c)
	})
}
