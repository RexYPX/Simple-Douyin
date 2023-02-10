package rpc

import (
	"context"
	"time"

	"Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/kitex_gen/publish/publishservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// userid to videolist
func PublishList(ctx context.Context, req *publish.PublishListRequest) (*publish.PublishListResponse, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, err
	}

	return resp, nil
}
