package service

import (
	"context"

	"github.com/DontSerious/DouSheng/cmd/relation/dal/db"
	"github.com/DontSerious/DouSheng/cmd/relation/pack"
	"github.com/DontSerious/DouSheng/kitex_gen/relation"
)

type QueryFollowerService struct {
	ctx context.Context
}

func NewQueryFollowerService(ctx context.Context) *QueryFollowerService {
	return &QueryFollowerService {
		ctx: ctx,
	}
}

func (s *QueryFollowerService) QueryFollower(req *relation.QueryFollowerRequest) ([]int64, error) {
	res, err := db.QueryFollower(s.ctx, req.UserId)
	userIds := make([]int64, len(res))

	if err != nil {
		return nil, err
	}

	rales := pack.Relas(res)
	for i := 0; i < len(rales); i++ {
		userIds[i] = rales[i].FollowerId
	}

	return userIds, nil
}