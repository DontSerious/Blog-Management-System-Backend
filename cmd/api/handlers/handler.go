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
		"status_code": 	statusCode,
		"status_msg":  	statusMsg,
		"data": 		data,
	})
}