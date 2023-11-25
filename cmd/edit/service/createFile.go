package service

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
	"os"
)

type CreateFileService struct {
	ctx context.Context
}

func NewCreateFileService(ctx context.Context) *CreateFileService {
	return &CreateFileService{
		ctx: ctx,
	}
}

func (s *CreateFileService) CreateFile(req *edit.CreateFileRequest) (statusCode int64, err error) {
	path := constants.EditDirectory + req.Path
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.WriteFile(path, nil, 0644); err != nil {
			return errno.ServiceErrCode, err
		}
	} else if err != nil {
		return errno.ServiceErrCode, err
	}

	return errno.SuccessCode, nil
}
