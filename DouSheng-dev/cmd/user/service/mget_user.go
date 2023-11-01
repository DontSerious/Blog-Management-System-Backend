package service

import (
	"context"

	"github.com/DontSerious/DouSheng/cmd/user/dal/db"
	"github.com/DontSerious/DouSheng/cmd/user/pack"
	"github.com/DontSerious/DouSheng/cmd/user/rpc"
	"github.com/DontSerious/DouSheng/kitex_gen/relation"
	"github.com/DontSerious/DouSheng/kitex_gen/user"
	"github.com/DontSerious/DouSheng/pkg/errno"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// 获取用户（列表）信息
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) (users []*user.User, statusCode int64, err error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	users = pack.Users(modelUsers)
	//is_follow调用relation服务
	followIds, statusCode, err := rpc.QueryFollow(context.Background(), &relation.QueryFollowRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, statusCode, err
	}
	for _, u := range users {
		if _, ok := followIds[u.Id]; ok {
			u.IsFollow = true
		}
	}
	return users, errno.SuccessCode, nil
}
