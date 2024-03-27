package handlers

import (
	"context"
	"io"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/api/rpc"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type PathParam struct {
	Path string `json:"path" form:"path"`
}

// 保存文件参数
type FileParam struct {
	Path    string `json:"path" form:"path"`
	Content string `json:"content" form:"content"`
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

func SaveFile(c *gin.Context) {
	var fileParam FileParam

	err := c.ShouldBind(&fileParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.SaveFile(context.Background(), &edit.SaveFileRequest{
		Path:    fileParam.Path,
		Content: fileParam.Content,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}

func DelAll(c *gin.Context) {
	var pathParam PathParam

	err := c.ShouldBind(&pathParam)
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	resp, err := rpc.DelAll(context.Background(), &edit.DelAllRequest{
		Path: pathParam.Path,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
		return
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}

func UploadFile(c *gin.Context) {
	var resp *edit.UploadFileResponse
	var err error

	// 获取上传的文件信息
	path := c.PostForm("path")
	fileSet, err := c.MultipartForm()
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return
	}

	// 打开文件
	for _, file := range fileSet.File["file"] {
		fs, err := file.Open()
		if err != nil {
			SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
			return
		}
		defer fs.Close()

		// 读取文件内容
		fileContent, err := io.ReadAll(fs)
		if err != nil {
			SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
			return
		}

		path = path + "/" + file.Filename

		// 发送请求
		resp, err = rpc.UploadFile(context.Background(), &edit.UploadFileRequest{
			File: fileContent,
			Path: path,
		})
		if err != nil {
			SendErrResponse(c, errno.ParamErrCode, errno.ParamErr.ErrMsg)
			return
		}
	}

	SendSuccResponse(c, resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg, nil)
}
