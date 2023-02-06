package service

import (
	"Simple-Douyin/cmd/message/dal/db"
	"Simple-Douyin/kitex_gen/message"
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type SendMessageService struct {
	ctx context.Context
}

func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

type Message struct {
	gorm.Model
	MessageId  int64  `grom:"primaryKey;autoIncrement" json:"id"`
	UserId     int64  `gorm:"index:idx_member, priority:1, not null" json:"user_id"`
	ToUserId   int64  `gorm:"index:idx_member, priority:2, not null" json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime string `gorm:"index;autoUpdateTime:nano" json:"create_time"`
}

func (s *SendMessageService) SendMessage(req *message.MessageActionRequest) error {
	if req.ActionType != 1 {
		return errors.New("msg: invaild action type")
	}

	klog.Debug("SendMessage: ", req)

	// TODO token -> userid
	userid := int64(1)

	msg := &db.Message{
		UserId:     userid,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		CreateTime: time.Now().Unix(),
	}

	if err := db.CreateMessage(s.ctx, msg); err != nil {
		klog.Info(err.Error())
	}

	return nil
}
