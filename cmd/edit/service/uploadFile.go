package service

import (
	"context"
	"os"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type UploadFileService struct {
	ctx context.Context
}

func NewUploadFileService(ctx context.Context) *UploadFileService {
	return &UploadFileService{
		ctx: ctx,
	}
}

func (s *UploadFileService) UploadFile(req *edit.UploadFileRequest) (statusCode int64, err error) {
	path := constants.EditDirectory + req.Path

	err = os.WriteFile(path, req.File, 0644)
	if err != nil {
		return errno.ServiceErr.ErrCode, err
	}

	return errno.SuccessCode, nil
}
