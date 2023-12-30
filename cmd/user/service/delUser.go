package service

import (
	"Bishe/be/cmd/user/dal/db"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
)

type DelUserService struct {
	ctx context.Context
}

func NewDelUserService(ctx context.Context) *DelUserService {
	return &DelUserService{
		ctx: ctx,
	}
}

func (s *DelUserService) DelUser(req *user.DelUserRequest) (statusCode int64, err error) {
	err = db.DelUser(s.ctx, req.UserId)
	if err != nil {
		return errno.UserNotExistErrCode, err
	}

	return errno.SuccessCode, nil
}
