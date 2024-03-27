package main

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/pack"
	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/service"
	user "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
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

// ChangePWD implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePWD(ctx context.Context, req *user.ChangePWDRequest) (resp *user.ChangePWDResponse, err error) {
	resp = new(user.ChangePWDResponse)

	//检查参数是否合法
	if len(req.UserId) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewChangePWDService(ctx).ChangePWD(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(statusCode, "修改密码成功")
	return resp, nil
}

// DelUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DelUser(ctx context.Context, req *user.DelUserRequest) (resp *user.DelUserResponse, err error) {
	resp = new(user.DelUserResponse)

	//检查参数是否合法
	if len(req.UserId) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.ParamErr.ErrMsg)
		return resp, nil
	}

	statusCode, err := service.NewDelUserService(ctx).DelUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(statusCode, "删除用户成功")
	return resp, nil
}

// GetAllUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAllUser(ctx context.Context) (resp *user.GetAllUserResponse, err error) {
	resp = new(user.GetAllUserResponse)

	users, statusCode, err := service.NewGetAllUserService(ctx).GetAllUser()
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.UserList = pack.BuildUserList(users)
	resp.BaseResp = pack.BuildBaseResponse(statusCode, "获取全部用户成功")

	return resp, nil
}
