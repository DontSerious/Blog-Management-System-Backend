package rpc

import (
	"Bishe/be/kitex_gen/userInfo"
	"Bishe/be/kitex_gen/userInfo/userinfoservice"
	"Bishe/be/pkg/constants"
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userInfoClient userinfoservice.Client

func initUserInfoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userinfoservice.NewClient(
		constants.UserInfoServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}

	userInfoClient = c
}

func QueryUserInfo(ctx context.Context, req *userinfo.GetUserInfoRequest) (*userinfo.GetUserInfoResponse, error) {
	return userInfoClient.GetUserInfo(ctx, req)
}

func UpdateUserInfo(ctx context.Context, req *userinfo.SetUserInfoRequest) (*userinfo.SetUserInfoResponse, error) {
	return userInfoClient.SetUserInfo(ctx, req)
}