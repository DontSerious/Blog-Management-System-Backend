package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// 检查用户名和密码是否正确，正确返回idStr
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (idStr string, statusCode int64, err error) {
	// 判断用户是否存在
	user, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return "", errno.UserNotExistErrCode, errno.UserNotExistErr
	}

	// 验证密码
	if user.Password != req.Password {
		return "", errno.LoginErrCode, errno.LoginErr
	}

	return user.ObjectID.Hex(), errno.SuccessCode, nil
}
