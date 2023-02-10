package rpc

import (
	"context"
	"time"

	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/favorite/favoriteservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
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
	favoriteClient = c
}

// FavoriteList 此接口用于从 favorite 获取 FavoriteList 的信息
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, err
	}

	return resp, nil
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
