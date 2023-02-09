package service

import (
	"Simple-Douyin/cmd/message/dal/db"
	"Simple-Douyin/kitex_gen/message"
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type SendMessageService struct {
	ctx context.Context
}

func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

func (s *SendMessageService) SendMessage(req *message.MessageActionRequest) error {
	if req.ActionType != 1 {
		return errors.New("msg: invaild action type")
	}

	msg := &db.Message{
		UserId:     req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		CreateTime: time.Now().Unix(),
	}

	if err := db.CreateMessage(s.ctx, msg); err != nil {
		klog.Info(err.Error())
	}

	return nil
}
