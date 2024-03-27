package service

import (
	"context"
	"os"

	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/errno"
)

type SaveFileService struct {
	ctx context.Context
}

func NewSaveFileService(ctx context.Context) *SaveFileService {
	return &SaveFileService{
		ctx: ctx,
	}
}

func (s *SaveFileService) SaveFile(req *edit.SaveFileRequest) (statusCode int64, err error) {
	path := constants.EditDirectory + req.Path

	// 以写入模式打开文件，不创建文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errno.ServiceErr.ErrCode, err
	}
	defer file.Close()

	// 将新内容写入文件
	_, err = file.WriteString(req.Content)
	if err != nil {
		return errno.ServiceErr.ErrCode, err
	}

	return errno.SuccessCode, nil
}
