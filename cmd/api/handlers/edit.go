package handlers

import (
	"Bishe/be/cmd/api/rpc"
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

type PathParam struct {
	Path string `json:"path" form:"path"`
}

func GetDirTree(c *gin.Context) {
	resp, err := rpc.GetDirTree(context.Background())
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.DirTree)
}

func GetFile(c *gin.Context) {
	var pathParam PathParam

	err := c.ShouldBind(&pathParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.GetFile(context.Background(), &edit.GetFileRequest{
		Path: pathParam.Path,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, resp.FileContent)
}

func CreateFile(c *gin.Context) {
	var pathParam PathParam

	err := c.ShouldBind(&pathParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.CreateFile(context.Background(), &edit.CreateFileRequest{
		Path: pathParam.Path,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}

func CreateDir(c *gin.Context) {
	var pathParam PathParam

	err := c.ShouldBind(&pathParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.CreateDir(context.Background(), &edit.CreateDirRequest{
		Path: pathParam.Path,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}