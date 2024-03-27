package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type GetAllUserService struct {
	ctx context.Context
}

func NewGetAllUserService(ctx context.Context) *GetAllUserService {
	return &GetAllUserService{
		ctx: ctx,
	}
}

func (s *GetAllUserService) GetAllUser() (users []*db.User, statusCode int64, err error) {
	users, err = db.GetAllUser(s.ctx)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}

	return users, errno.SuccessCode, nil
}
