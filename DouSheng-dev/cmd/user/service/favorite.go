package service

import (
	"context"

	"github.com/DontSerious/DouSheng/kitex_gen/user"
	"github.com/DontSerious/DouSheng/cmd/user/dal/db"
	"github.com/DontSerious/DouSheng/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) Favorite(req *user.FavoriteOperationRequest) (statusCode int64, err error) {
	err = db.Favorite(s.ctx, req.UserId, req.VideoAuthor, req.ActionType)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	return errno.SuccessCode, nil
}
