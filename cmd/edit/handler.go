package main

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/edit/pack"
	"github.com/DontSerious/Blog-Management-System-Backend/cmd/edit/service"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

// EditServiceImpl implements the last service interface defined in the IDL.
type EditServiceImpl struct{}

// GetDirTree implements the EditServiceImpl interface.
func (s *EditServiceImpl) GetDirTree(ctx context.Context) (resp *edit.GetDirTreeResponse, err error) {
	resp = new(edit.GetDirTreeResponse)

	res, statusCode, err := service.NewGetDirTreeService(ctx).GetDirTree()
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取文件夹结构成功")
	resp.DirTree = res

	return resp, nil
}

// GetFile implements the EditServiceImpl interface.
func (s *EditServiceImpl) GetFile(ctx context.Context, req *edit.GetFileRequest) (resp *edit.GetFileResponse, err error) {
	resp = new(edit.GetFileResponse)

	if len(req.Path) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	res, statusCode, err := service.NewGetFileService(ctx).GetFile(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取文件内容成功")
	resp.FileContent = *res

	return resp, nil
}

// CreateFile implements the EditServiceImpl interface.
func (s *EditServiceImpl) CreateFile(ctx context.Context, req *edit.CreateFileRequest) (resp *edit.CreateFileResponse, err error) {
	resp = new(edit.CreateFileResponse)

	if len(req.Path) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewCreateFileService(ctx).CreateFile(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "创建文件成功")

	return resp, nil
}

// CreateDir implements the EditServiceImpl interface.
func (s *EditServiceImpl) CreateDir(ctx context.Context, req *edit.CreateDirRequest) (resp *edit.CreateDirResponse, err error) {
	resp = new(edit.CreateDirResponse)

	if len(req.Path) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewCreateDirService(ctx).CreateDir(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "创建文件夹成功")
	return
}

// SaveFile implements the EditServiceImpl interface.
func (s *EditServiceImpl) SaveFile(ctx context.Context, req *edit.SaveFileRequest) (resp *edit.SaveFileResponse, err error) {
	resp = new(edit.SaveFileResponse)

	if len(req.Path) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewSaveFileService(ctx).SaveFile(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "保存成功")

	return
}

// DelAll implements the EditServiceImpl interface.
func (s *EditServiceImpl) DelAll(ctx context.Context, req *edit.DelAllRequest) (resp *edit.DelAllResponse, err error) {
	resp = new(edit.DelAllResponse)

	if len(req.Path) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewDelAllService(ctx).DelAll(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "删除成功")
	return
}

// UploadFile implements the EditServiceImpl interface.
func (s *EditServiceImpl) UploadFile(ctx context.Context, req *edit.UploadFileRequest) (resp *edit.UploadFileResponse, err error) {
	resp = new(edit.UploadFileResponse)

	statusCode, err := service.NewUploadFileService(ctx).UploadFile(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(statusCode, "上传文件成功")

	return resp, nil
}
