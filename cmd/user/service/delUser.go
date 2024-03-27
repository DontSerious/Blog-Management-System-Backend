package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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
