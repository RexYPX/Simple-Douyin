package rpc

import (
	"Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/kitex_gen/publish/publishservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(constants.ApiServiceName),
	// 	provider.WithExportEndpoint(constants.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := publishservice.NewClient(
		constants.PublishServiceName,
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
	publishClient = c
}

// PublishAction create publish info
func PublishAction(ctx context.Context, req *publish.PublishActionRequest) error {
	resp, err := publishClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return nil
}

// PublishList query list of publish info
func PublishList(ctx context.Context, req *publish.PublishListRequest) (*publish.PublishListResponse, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return resp, nil
}
