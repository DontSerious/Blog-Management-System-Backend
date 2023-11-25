package main

import (
	edit "Bishe/be/kitex_gen/edit/editservice"
	"Bishe/be/pkg/constants"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.EditResolveTCPAddr)
	if err != nil {
		panic(err)
	}

	svr := edit.NewServer(
		new(EditServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.EditServiceName}),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
