package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/dal/db"
	userInfo "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/userInfo"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type QueryUserInfoService struct {
	ctx context.Context
}

func NewQueryUserInfoService(ctx context.Context) *QueryUserInfoService {
	return &QueryUserInfoService{
		ctx: ctx,
	}
}

func (s *QueryUserInfoService) QueryUserInfo(req *userInfo.GetUserInfoRequest) (userInfo *db.UserInfo, statusCode int64, err error) {
	userInfo, err = db.QueryUserInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}

	return userInfo, errno.SuccessCode, nil
}
