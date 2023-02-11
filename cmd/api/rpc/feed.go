package rpc

import (
	"Simple-Douyin/kitex_gen/feed"
	"Simple-Douyin/kitex_gen/feed/feedservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(constants.ApiServiceName),
	// 	provider.WithExportEndpoint(constants.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := feedservice.NewClient(
		constants.FeedServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

// Feed feed videos to user
func Feed(ctx context.Context, req *feed.FeedRequest) (int64, []*feed.Video, error) {
	log.Println("[ypx debug] api rpc Feed: prepare to feedClient.Feed(ctx, req)")
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		log.Println("[ypx debug] api rpc Feed: feedClient.Feed(ctx, req) err ", err)
		return 0, nil, err
	}
	if resp.StatusCode != 0 {
		log.Println("[ypx debug] api rpc Feed: feedClient.Feed(ctx, req) resp.StatusCode != 0", resp.StatusCode, " ", resp.StatusMsg)
		return 0, nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	log.Println("[ypx debug] api rpc Feed: feedClient.Feed(ctx, req) success")
	return resp.NextTime, resp.VideoList, nil
}
