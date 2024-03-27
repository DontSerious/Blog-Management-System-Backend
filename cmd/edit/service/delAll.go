package service

import (
	"context"
	"os"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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
