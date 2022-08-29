package middlewares

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Download 文件下载
func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if strings.HasPrefix(uri, "/upload") && c.Request.Method == "GET" {
			bytes, err := ioutil.ReadFile("." + uri)
			if err != nil {
				c.Status(http.StatusNotFound)
				c.Abort()
				return
			}
			// uriArr := strings.Split(uri, "/")
			// fileName := uriArr[len(uriArr)-1]
			// // 判断是否为文件
			// if len(strings.Split(fileName, ".")) < 2 {
			// 	c.Next()
			// 	return
			// }
			// fileNameArr := strings.Split(fileName, ".")
			// fileExt := fileNameArr[len(fileNameArr)-1]
			// switch fileExt {
			// case "jpeg":
			// 	fallthrough
			// case "jpg":
			// 	c.Writer.Header().Add("Content-Type", "image/jpg")
			// case "png":
			// 	c.Writer.Header().Add("Content-Type", "image/png")
			// case "js":
			// 	c.Writer.Header().Add("Content-Type", "image/git")
			// }
			c.Writer.Write(bytes)
			c.Abort()
			return
		}
	}
}
