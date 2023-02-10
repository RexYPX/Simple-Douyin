package rpc

import (
	"context"
	"time"

	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/comment/commentservice"
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}

// FavoriteCount 此接口用于从 favorite 获取 FavoriteList 的信息
func CommentCount(ctx context.Context, req *comment.CommentListRequest) (int64, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, err
	}

	return int64(len(resp.CommentList)), nil
}

// FavoriteCount 此接口用于从 favorite 获取 FavoriteCount 的信息
func FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	resp, err := favoriteClient.FavoriteCount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// IsFavorite 此接口用于从 favorite 获取 IsFavorite
func IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	resp, err := favoriteClient.IsFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
