package handlers

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/api/rpc"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type ChangePWDParam struct {
	UserId   string `json:"_id" form:"_id"`
	Password string `json:"password" form:"password"`
}

func ChangePWD(c *gin.Context) {
	var p ChangePWDParam

	err := c.ShouldBind(&p)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.ChangePWD(context.Background(), &user.ChangePWDRequest{
		UserId:   p.UserId,
		Password: p.Password,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}

func DelUser(c *gin.Context) {
	userId := c.Query("_id")

	resp, err := rpc.DelUser(context.Background(), &user.DelUserRequest{
		UserId: userId,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}

func GetAllUser(c *gin.Context) {
	resp, err := rpc.GetAllUser(context.Background())
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.UserList)
}
