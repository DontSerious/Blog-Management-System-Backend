package handlers

import (
	"Bishe/be/cmd/api/rpc"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	var RegisterParam RegisterParam

	// 获取参数
	RegisterParam.UserName = c.Query("username")
	RegisterParam.PassWord = c.Query("password")

	//将注册信息写入数据库
	user_id, statusCode, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: RegisterParam.UserName,
		Password: RegisterParam.PassWord,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "注册成功",
		"user_id":     user_id,
	})
}