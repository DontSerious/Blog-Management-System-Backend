package rpc

import (
	"context"
	"time"

	"github.com/DontSerious/DouSheng/cmd/relation/pack"
	"github.com/DontSerious/DouSheng/kitex_gen/relation"
	"github.com/DontSerious/DouSheng/kitex_gen/user/userservice"
	"github.com/DontSerious/DouSheng/pkg/constants"
	"github.com/DontSerious/DouSheng/pkg/errno"
	"github.com/DontSerious/DouSheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

// user服务发现
func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// 发现user的MGetUser
func MGetUser(ctx context.Context, req *relation.QueryUserListRequest) ([]*relation.User, error) {
	// 打包为user.MGetUserRequest
	r := pack.MGetUserReq(req.UserId, req.UserIds)

	resp, err := userClient.MGetUser(ctx, r)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	userList := pack.UserList(resp.UserList)
	return userList, nil
}
