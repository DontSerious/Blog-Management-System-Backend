package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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
