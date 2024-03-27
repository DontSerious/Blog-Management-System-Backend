package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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
