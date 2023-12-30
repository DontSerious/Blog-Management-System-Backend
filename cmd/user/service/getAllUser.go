package service

import (
	"Bishe/be/cmd/user/dal/db"
	"Bishe/be/pkg/errno"
	"context"
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
