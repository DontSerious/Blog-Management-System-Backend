package service

import (
	"Bishe/be/cmd/userInfo/dal/db"
	userInfo "Bishe/be/kitex_gen/userInfo"
	"Bishe/be/pkg/errno"
	"context"
)

type UpdateUserInfoService struct {
	ctx context.Context
}

func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{
		ctx: ctx,
	}
}

func (s *UpdateUserInfoService) UpdateUserInfo(req *userInfo.SetUserInfoRequest) (statusCode int64, err error) {
	// 判断用户是否存在
	_, err = db.QueryUserInfo(s.ctx, req.UserId)
	if err != nil {
		return errno.UserNameHasUsedErrCode, errno.UserNameHasUsedErr
	}

	err = db.UpdateUserInfo(s.ctx, req.UserId, &db.UserInfo{
		Categories: req.UserInfo.Categories,
		Tags: req.UserInfo.Tags,
	})
	if err != nil {
		return errno.ServiceErrCode, err
	}

	return errno.SuccessCode, nil
}