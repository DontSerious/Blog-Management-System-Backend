package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrResponse(c *gin.Context, statusCode int64, errMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  errMsg,
	})
}

func SendSuccResponse(c *gin.Context, statusCode int64, statusMsg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"data":        data,
	})
}

// func SendFileResponse(c *gin.Context, statusCode int64, statusMsg string, fileName string, data []byte) {
// 	// 设置响应头，指定传输的文件类型
// 	c.Header("Content-Type", "application/octet-stream")
// 	// 设置 Content-Disposition 头部，指定文件名和下载方式
// 	c.Header("Content-Disposition", "attachment; filename=aaa.txt")
// 	c.Header("Content-Length", strconv.FormatInt(data.Length, 10)
// 	c.Writer.Write(data)
// 	c.File()
// }
