package service

import (
	"context"
	"os"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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
