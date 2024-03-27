package rpc

import (
	"context"
	"time"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user/userservice"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRPC() {
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

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return userClient.CreateUser(ctx, req)
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (*user.CheckUserResponse, error) {
	return userClient.CheckUser(ctx, req)
}

func ChangePWD(ctx context.Context, req *user.ChangePWDRequest) (*user.ChangePWDResponse, error) {
	return userClient.ChangePWD(ctx, req)
}

func DelUser(ctx context.Context, req *user.DelUserRequest) (*user.DelUserResponse, error) {
	return userClient.DelUser(ctx, req)
}

func GetAllUser(ctx context.Context) (*user.GetAllUserResponse, error) {
	return userClient.GetAllUser(ctx)
}
