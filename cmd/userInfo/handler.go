package main

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/pack"
	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/service"
	userinfo "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/userInfo"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *userinfo.GetUserInfoRequest) (resp *userinfo.GetUserInfoResponse, err error) {
	resp = new(userinfo.GetUserInfoResponse)

	if len(req.UserId) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	res, statusCode, err := service.NewQueryUserInfoService(ctx).QueryUserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取用户信息成功")
	resp.UserInfo = pack.UserInfo(res)

	return resp, nil
}

// SetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetUserInfo(ctx context.Context, req *userinfo.SetUserInfoRequest) (resp *userinfo.SetUserInfoResponse, err error) {
	resp = new(userinfo.SetUserInfoResponse)

	if len(req.UserId) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewUpdateUserInfoService(ctx).UpdateUserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "更新用户信息成功")

	return resp, nil
}
