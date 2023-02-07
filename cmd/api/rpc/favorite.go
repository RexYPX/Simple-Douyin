package rpc

import (
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/favorite/favoriteservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	fmt.Println("resp.StatusMsg:", resp.StatusMsg)
	return nil
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp.VideoList, nil
}
