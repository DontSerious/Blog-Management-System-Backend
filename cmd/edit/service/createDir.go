package service

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
	"os"
)

type CreateDirService struct {
	ctx context.Context
}

func NewCreateDirService(ctx context.Context) *CreateDirService {
	return &CreateDirService{
		ctx: ctx,
	}
}

func (s *CreateDirService) CreateDir(req *edit.CreateDirRequest) (statusCode int64, err error) {
	path := constants.EditDirectory + req.Path

	// 创建文件夹
	if err := os.MkdirAll(path, 0755); err != nil {
		return errno.ServiceErrCode, err
	}

	return errno.SuccessCode, nil
}
