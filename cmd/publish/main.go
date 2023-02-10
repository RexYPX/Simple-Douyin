package main

import (
	"Simple-Douyin/cmd/publish/dal"
	"Simple-Douyin/cmd/publish/rpc"
	publish "Simple-Douyin/kitex_gen/publish/publishservice"
	"Simple-Douyin/pkg/constants"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	rpc.InitRPC()
	dal.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(constants.TCP, constants.PublishServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(constants.CommentServiceName),
	//  provider.WithExportEndpoint(constants.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	// svr := comment.NewServer(new(CommentServiceImpl),
	// 	server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
	// 	server.WithMuxTransport(),
	// 	server.WithMiddleware(mw.CommonMiddleware),
	// 	server.WithMiddleware(mw.ServerMiddleware),
	// 	server.WithSuite(tracing.NewServerSuite()),
	// )
	// err = svr.Run()
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.PublishServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
