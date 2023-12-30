package service

import (
	"Bishe/be/cmd/user/dal/db"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
)

type ChangePWDService struct {
	ctx context.Context
}

func NewChangePWDService(ctx context.Context) *ChangePWDService {
	return &ChangePWDService{
		ctx: ctx,
	}
}

func (s *ChangePWDService) ChangePWD(req *user.ChangePWDRequest) (statusCode int64, err error) {
	err = db.ChangePWD(s.ctx, req.UserId, req.Password)
	if err != nil {
		return errno.UserNotExistErrCode, err
	}

	return errno.SuccessCode, nil
}
