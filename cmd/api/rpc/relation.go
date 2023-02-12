package rpc

import (
	"Simple-Douyin/cmd/api/biz/handler/pack"
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/relation"
	"Simple-Douyin/kitex_gen/relation/relationservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/mw"
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(constants.ApiServiceName),
	// 	provider.WithExportEndpoint(constants.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = c

	log.Println("[ypx debug] initRelation success")
}

// RelationAction create/delete relation info
func RelationAction(ctx context.Context, req *relation.RelationActionRequest, resp *api.RelationActionResponse) error {
	rpcResp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if rpcResp.StatusCode != 0 {
		return err
	}

	pack.RelationAction2ApiAction(rpcResp, resp)

	return nil
}

// RelationFollowList query list of follows' info
func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest, resp *api.RelationFollowListResponse) error {
	rpcResp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return err
	}
	if rpcResp.StatusCode != 0 {
		return err
	}

	pack.RelationFollowList2ApiFollowList(rpcResp, resp)

	return nil
}

// RelationFollowerList query list of followers' info
func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest, resp *api.RelationFollowerListResponse) error {
	rpcResp, err := relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return err
	}
	if rpcResp.StatusCode != 0 {
		return err
	}

	pack.RelationFollowerList2ApiFollowerList(rpcResp, resp)

	return nil
}

// RelationFriendList query list of friends' info
func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest, resp *api.RelationFriendListResponse) error {
	rpcResp, err := relationClient.RelationFriendList(ctx, req)
	if err != nil {
		return err
	}
	if rpcResp.StatusCode != 0 {
		return err
	}

	pack.RelationFriendList2ApiFriendList(rpcResp, resp)

	return nil
}
