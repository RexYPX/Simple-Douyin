package main

import (
	relation "Simple-Douyin/kitex_gen/relation/relationservice"
	"Simple-Douyin/pkg/constants"
	"log"
	"net"

	"Simple-Douyin/cmd/relation/dal"
	"Simple-Douyin/cmd/relation/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	rpc.Init()
	dal.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr(constants.TCP, constants.RelationServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationServiceName}),
	)

	err = svr.Run()

	// svr := relation.NewServer(new(RelationServiceImpl))
	// err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
