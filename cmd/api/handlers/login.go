package handlers

import (
	"Bishe/be/cmd/api/rpc"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"

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
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	//将注册信息写入数据库
	resp, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: loginParam.UserName,
		Password: loginParam.PassWord,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.UserId)
}