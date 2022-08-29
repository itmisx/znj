package routers

import (
	"znj/controllers"

	"github.com/gin-gonic/gin"
)

// UploadRouter 上传路由
func UploadRouter(r *gin.RouterGroup) {
	upload := controllers.Upload{}
	// vditor上传
	r.POST("/vditorUpload", func(ctx *gin.Context) {
		upload.VditorUpload(ctx)
	})
	// wangeditor上传
	r.POST("/weditorUpload", func(ctx *gin.Context) {
		upload.WEditorUpload(ctx)
	})
}
