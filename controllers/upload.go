package controllers

import (
	"strings"

	"znj/logics"

	"github.com/gin-gonic/gin"
	"github.com/itmisx/i18n"
)

// Upload 上传
type Upload struct{}

// VditorUpload vditor上传
func (up *Upload) VditorUpload(c *gin.Context) {
	msg := ""

	upload := logics.Upload{}
	filepaths, filenames, errCode := upload.Upload(c)

	succMap := make(map[string]string) // 初始化，不初始化不能使用

	for index, filepath := range filepaths {
		succMap[filenames[index]] = filepath
	}
	lang := strings.Split(c.GetHeader("Accept-Language"), ",")[0]
	if errCode > 0 {
		msg = i18n.T(lang, errCode)
	}

	c.JSON(200, gin.H{
		"code": errCode,
		"data": gin.H{
			"succMap": succMap,
		},
		"msg": msg,
	})
}

// WEditorUpload weditor上传
func (up *Upload) WEditorUpload(c *gin.Context) {
	msg := ""

	upload := logics.Upload{}
	filepaths, _, errCode := upload.Upload(c)
	lang := strings.Split(c.GetHeader("Accept-Language"), ",")[0]
	if errCode > 0 {
		msg = i18n.T(lang, errCode)
	}

	c.JSON(200, gin.H{
		"errno": errCode,
		"data":  filepaths,
		"msg":   msg,
	})
}
