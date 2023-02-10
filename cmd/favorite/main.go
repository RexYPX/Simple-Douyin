package main

import (
	"Simple-Douyin/cmd/favorite/dal"
	favorite "Simple-Douyin/kitex_gen/favorite/favoriteservice"
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
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	if err != nil {
		panic(err)
	}

	Init()

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FavoriteServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
