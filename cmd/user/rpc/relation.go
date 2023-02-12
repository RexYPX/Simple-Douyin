package rpc

import (
	"Simple-Douyin/kitex_gen/relation"
	"Simple-Douyin/kitex_gen/relation/relationservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = c
}

// IsFollow 此接口用于判断id1是否关注id2
func IsFollow(ctx context.Context, req *relation.RelationIsFollowRequest) (bool, error) {
	resp, err := relationClient.RelationIsFollow(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 0 {
		return false, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp.IsFollow, nil
}

func FollowerCount(ctx context.Context, req *relation.RelationFollowerCountRequest) (int64, error) {
	resp, err := relationClient.RelationFollowerCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != 0 {
		return -1, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp.FollowerCount, nil
}

func FollowCount(ctx context.Context, req *relation.RelationFollowCountRequest) (int64, error) {
	resp, err := relationClient.RelationFollowCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != 0 {
		return -1, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp.FollowCount, nil
}
