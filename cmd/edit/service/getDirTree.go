package service

import (
	"context"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/edit/utils"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type GetDirTreeService struct {
	ctx context.Context
}

func NewGetDirTreeService(ctx context.Context) *GetDirTreeService {
	return &GetDirTreeService{
		ctx: ctx,
	}
}

func (s *GetDirTreeService) GetDirTree() (dirTree []*edit.DataNode, statusCode int64, err error) {
	fileTree, err := utils.ReadDirectory(constants.EditDirectory)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}

	return fileTree.Children, errno.SuccessCode, nil
}
