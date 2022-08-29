//Package logics 上传文件逻辑
package logics

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Upload 上传
type Upload struct {
}

//Upload 上传
func (up *Upload) Upload(c *gin.Context) (urls []string, filenames []string, errCode int64) {
	//判断目标文件夹是否存在
	year, month, _ := time.Now().Date()
	monthDir := fmt.Sprintf("%d%d", year, month)
	saveDir := "./upload/" + monthDir + "/" + strconv.Itoa(time.Now().Day())
	_, err2 := os.Stat(saveDir)
	//不存在就创建
	if err2 != nil {
		err3 := os.MkdirAll(saveDir, os.ModePerm)
		if err3 != nil {
			return
		}
	}
	// 接收上传内容
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	files := form.File["file[]"]
	if files == nil {
		file, err1 := c.FormFile("file")
		if err1 == nil {
			//生成随机文件名
			fileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(int64(rand.Int31()), 10)
			//文件后置
			ext := filepath.Ext(file.Filename)
			// 保存文件
			dst := saveDir + "/" + fileName + ext
			err4 := c.SaveUploadedFile(file, dst)
			if err4 == nil {
				errCode = 0
				filenames = append(filenames, file.Filename)
				urls = append(urls, dst)
			}
		}

	} else {
		for _, file := range files {
			//生成随机文件名
			fileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(int64(rand.Int31()), 10)
			//文件后置
			ext := filepath.Ext(file.Filename)
			// 保存文件
			dst := saveDir + "/" + fileName + ext
			err4 := c.SaveUploadedFile(file, dst)
			if err4 == nil {
				errCode = 0
				filenames = append(filenames, file.Filename)
				urls = append(urls, dst)
			}
		}
	}
	return
}
