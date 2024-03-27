package main

import (
	"log"
	"net"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/dal"
	userinfo "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/userInfo/userinfoservice"
	"github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserInfoResolveTCPAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := userinfo.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserInfoServiceName}),
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
