package service

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
	"os"
)

type DelAllService struct {
	ctx context.Context
}

func NewDelAllService(ctx context.Context) *DelAllService {
	return &DelAllService{
		ctx: ctx,
	}
}

func (s *DelAllService) DelAll(req *edit.DelAllRequest) (statusCode int64, err error) {
	path := constants.EditDirectory + req.Path

	err = os.RemoveAll(path)
	if err != nil {
		return errno.ServiceErr.ErrCode, err
	}

	return errno.SuccessCode, nil
}
