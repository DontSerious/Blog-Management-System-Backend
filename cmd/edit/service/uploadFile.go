package service

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
	"os"
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
