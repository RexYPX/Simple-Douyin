package main

import (
	"Simple-Douyin/cmd/message/service"
	message "Simple-Douyin/kitex_gen/message"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

// MeassgeServiceImpl implements the last service interface defined in the IDL.
type MeassgeServiceImpl struct{}

// MessageAction implements the MeassgeServiceImpl interface.
func (s *MeassgeServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp = new(message.MessageActionResponse)

	err = service.NewSendMessageService(ctx).SendMessage(req)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "message send failed."

		return resp, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = "message send success!"

	klog.Info(resp)

	return resp, nil
}

// MessageChat implements the MeassgeServiceImpl interface.
func (s *MeassgeServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	resp = new(message.MessageChatResponse)

	msgs, err := service.NewGetMessageHistoryService(ctx).GetMessageHistory(req)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "get message history failed"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "get message history success"
	resp.MessageList = msgs

	return resp, nil
}
