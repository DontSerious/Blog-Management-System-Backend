package handlers

import (
	"Bishe/be/cmd/api/rpc"
	userinfo "Bishe/be/kitex_gen/userInfo"
	"Bishe/be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	UserId string `json:"_id" form:"_id"`
}

type UpdateParam struct {
	UserId     string   `json:"_id" form:"_id"`
	Categories []string `json:"Categories" form:"Categories"`
	Tags       []string `json:"Tags" form:"Tags"`
}

func Query(c *gin.Context) {
	var queryParam QueryParam

	err := c.ShouldBind(&queryParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.QueryUserInfo(context.Background(), &userinfo.GetUserInfoRequest{
		UserId: queryParam.UserId,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.UserInfo)
}

func Update(c *gin.Context) {
	var updateParam UpdateParam

	err := c.BindJSON(&updateParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.UpdateUserInfo(context.Background(), &userinfo.SetUserInfoRequest{
		UserId: updateParam.UserId,
		UserInfo: &userinfo.UserInfo{
			Categories: updateParam.Categories,
			Tags:       updateParam.Tags,
		},
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}
