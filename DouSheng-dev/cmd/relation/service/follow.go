package service

import (
	"context"

	"github.com/DontSerious/DouSheng/cmd/relation/dal/db"
	"github.com/DontSerious/DouSheng/kitex_gen/relation"
	"github.com/DontSerious/DouSheng/pkg/errno"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{
		ctx: ctx,
	}
}

func (s *FollowService) Follow(req *relation.FollowRequest) error {
	var err error

	r := db.Relation{
		Follow_id:   req.ToUserId,
		Follower_id: req.UserId,
	}

	// 关注
	if req.ActionType == 1 {
		err = db.Follow(s.ctx, &r)
	} else if req.ActionType == 2 {
		// 取关
		err = db.CancelFollow(s.ctx, &r)
	} else {
		err = errno.ParamErr
	}

	return err
}
