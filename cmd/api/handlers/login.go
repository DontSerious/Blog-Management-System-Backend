package handlers

import (
	"Bishe/be/cmd/api/rpc"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var loginParam LoginParam

	// 获取参数
	err := c.ShouldBind(&loginParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, err)
		return
	}

	//将注册信息写入数据库
	user_id, statusCode, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: loginParam.UserName,
		Password: loginParam.PassWord,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "登录成功",
		"user_id":     user_id,
	})
}