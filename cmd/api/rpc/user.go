package rpc

import (
	"Bishe/be/kitex_gen/user"
	"Bishe/be/kitex_gen/user/userservice"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (user_id string, statusCode int64, err error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return "", errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return "", resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.SuccessCode, nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (user_id string, statusCode int64, err error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return "", errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return "", resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.SuccessCode, nil
}