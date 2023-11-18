package service

import (
	"Bishe/be/cmd/userInfo/dal/db"
	userInfo "Bishe/be/kitex_gen/userInfo"
	"Bishe/be/pkg/errno"
	"context"
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
	userInfo, err = db.QueryUserInfo(s.ctx, req.UserId);
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}

	return userInfo, errno.SuccessCode, nil
}