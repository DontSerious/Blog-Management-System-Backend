package service

import (
	"Bishe/be/cmd/user/dal/db"
	"Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{
		ctx: ctx,
	}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) (idStr string, statusCode int64, err error) {
	// 判断用户是否存在
	_, err = db.QueryUser(s.ctx, req.Username)
	if err == nil {
		return "", errno.UserNameHasUsedErrCode, errno.UserNameHasUsedErr
	}
	
	// 创建用户
	idStr, err = db.CreateUser(s.ctx, &db.User{
		UserName: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return "", errno.ServiceErrCode, err
	}

	return idStr, errno.SuccessCode, nil
}