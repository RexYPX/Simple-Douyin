package main

import (
	"Simple-Douyin/cmd/message/dal"
	message "Simple-Douyin/kitex_gen/message/meassgeservice"
	"Simple-Douyin/pkg/constants"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	klog.SetLevel(klog.LevelDebug)

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8801")
	if err != nil {
		panic(err)
	}

	Init()

	svr := message.NewServer(new(MeassgeServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MessageServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
