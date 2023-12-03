package rpc

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/kitex_gen/edit/editservice"
	"Bishe/be/pkg/constants"
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var editClient editservice.Client

func initEditRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := editservice.NewClient(
		constants.EditServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	editClient = c
}

func GetDirTree(ctx context.Context) (*edit.GetDirTreeResponse, error) {
	return editClient.GetDirTree(ctx)
}

func GetFile(ctx context.Context, req *edit.GetFileRequest) (*edit.GetFileResponse, error) {
	return editClient.GetFile(ctx, req)
}

func CreateFile(ctx context.Context, req *edit.CreateFileRequest) (*edit.CreateFileResponse, error) {
	return editClient.CreateFile(ctx, req)
}

func CreateDir(ctx context.Context, req *edit.CreateDirRequest) (*edit.CreateDirResponse, error) {
	return editClient.CreateDir(ctx, req)
}

func SaveFile(ctx context.Context, req *edit.SaveFileRequest) (*edit.SaveFileResponse, error) {
	return editClient.SaveFile(ctx, req)
}

func DelAll(ctx context.Context, req *edit.DelAllRequest) (*edit.DelAllResponse, error) {
	return editClient.DelAll(ctx, req)
}
