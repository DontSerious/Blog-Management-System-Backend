package service

import (
	"context"
	"os"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type GetFileService struct {
	ctx context.Context
}

func NewGetFileService(ctx context.Context) *GetFileService {
	return &GetFileService{
		ctx: ctx,
	}
}

func (s *GetFileService) GetFile(req *edit.GetFileRequest) (file *string, statusCode int64, err error) {
	path := constants.EditDirectory + req.Path

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, errno.ServiceErr.ErrCode, err
	}

	res := string(content)

	return &res, errno.SuccessCode, nil
}
