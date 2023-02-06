package rpc

import (
	"Simple-Douyin/kitex_gen/message"
	"Simple-Douyin/kitex_gen/message/meassgeservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"context"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient meassgeservice.Client

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := meassgeservice.NewClient(
		constants.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func SendMessage(ctx context.Context, req *message.MessageActionRequest) error {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return nil
}

func GetMessageHistory(ctx context.Context, req *message.MessageChatRequest) ([]*message.Message, error) {
	resp, err := messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp.MessageList, nil
}
