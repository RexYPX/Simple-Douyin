package main

//

import (
	user "Simple-Douyin/cmd/user/kitex_gen/user/userservice"
	"Simple-Douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"

	"Simple-Douyin/cmd/user/dal"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceAddr)
	if err != nil {
		panic(err)
	}

	dal.Init()

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
	)

	//svr := user.NewServer(new(UserServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
