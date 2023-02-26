package service

import (
	"Simple-Douyin/cmd/message/dal/db"
	"Simple-Douyin/cmd/message/pack"
	"Simple-Douyin/kitex_gen/message"
	"context"
)

type GetMessageHistoryService struct {
	ctx context.Context
}

func NewGetMessageHistoryService(ctx context.Context) *GetMessageHistoryService {
	return &GetMessageHistoryService{ctx: ctx}
}

func (s *GetMessageHistoryService) GetMessageHistory(req *message.MessageChatRequest) ([]*message.Message, error) {
	uid  := req.UserId
	tuid := req.ToUserId
    pst  := req.PreMsgTime

	dbmsgs, err := db.QueryMessageHistory(s.ctx, uid, tuid, pst)
	if err != nil {
		panic(err)
	}

	// db.Message -> message.Message
	msgs := pack.Messages(dbmsgs)

	return msgs, nil
}
