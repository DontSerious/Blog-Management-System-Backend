package service

import (
	"Bishe/be/cmd/edit/utils"
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"Bishe/be/pkg/errno"
	"context"
)

type GetDirTreeService struct {
	ctx context.Context
}

func NewGetDirTreeService(ctx context.Context) *GetDirTreeService {
	return &GetDirTreeService{
		ctx: ctx,
	}
}

func (s *GetDirTreeService) GetDirTree() (dirTree []*edit.FileNode, statusCode int64, err error) {
	fileTree, err := utils.ReadDirectory(constants.EditDirectory)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}

	return fileTree.Children, errno.SuccessCode, nil
}
