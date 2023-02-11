package rpc

import (
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/comment/commentservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(constants.ApiServiceName),
	// 	provider.WithExportEndpoint(constants.ExportEndpoint),
	// 	provider.WithInsecure(),
	// )
	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}

// CommentAction create/delete comment info
func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.Comment, error) {
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Comment, nil
}

// CommentList query list of comment info
func CommentList(ctx context.Context, req *comment.CommentListRequest) ([]*comment.Comment, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentList, nil
}
