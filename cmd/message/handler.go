package main

import (
	message "Simple-Douyin/cmd/message/kitex_gen/message"
	"context"
)

// MeassgeServiceImpl implements the last service interface defined in the IDL.
type MeassgeServiceImpl struct{}

// MessageAction implements the MeassgeServiceImpl interface.
func (s *MeassgeServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the MeassgeServiceImpl interface.
func (s *MeassgeServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
