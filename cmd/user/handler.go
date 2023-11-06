package main

import (
	"Bishe/be/cmd/user/pack"
	"Bishe/be/cmd/user/service"
	user "Bishe/be/kitex_gen/user"
	"Bishe/be/pkg/errno"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	//检查参数是否合法
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		resp.UserId = ""
		return resp, nil
	}

	//将用户名密码插入数据库，返回user_id
	idStr, statusCode, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserId = ""
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "注册用户成功")
	resp.UserId = idStr
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)

	//检查参数是否合法
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		resp.UserId = ""
		return resp, nil
	}

	//查询数据库，看用户名和密码是否正确，正确返回user_id
	idStr, statusCode, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserId = ""
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "用户名，密码正确")
	resp.UserId = idStr
	return resp, nil
}
