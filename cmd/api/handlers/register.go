package handlers

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/api/rpc"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type RegisterParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	var registerParam RegisterParam

	// 获取参数
	err := c.ShouldBind(&registerParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	//将注册信息写入数据库
	resp, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerParam.UserName,
		Password: registerParam.PassWord,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.UserId)
}
