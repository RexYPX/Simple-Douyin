package db

import (
	"Simple-Douyin/pkg/constants"
	"context"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	// MessageId  int64  `grom:"primaryKey;autoIncrement" json:"id"`
	UserId     int64  `gorm:"index:idx_member, priority:1, not null" json:"user_id"`
	ToUserId   int64  `gorm:"index:idx_member, priority:2, not null" json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `gorm:"index;autoUpdateTime:nano" json:"create_time"`
}

func (msg *Message) TableName() string {
	return constants.MessageTableName
}

func CreateMessage(ctx context.Context, msg *Message) error {
	return DB.WithContext(ctx).Create(msg).Error
}

// use to_user_id to queary message history.
func QueryMessageHistory(ctx context.Context, uid, tuid, pst int64) ([]*Message, error) {
	var resp []*Message

	if err := DB.WithContext(ctx).Model(&Message{}).Where("(user_id = ? and to_user_id = ? or user_id = ? and to_user_id = ?) and create_time > ?", uid, tuid, tuid, uid, pst).Find(&resp).Error; err != nil {
		return resp, err
	}

	return resp, nil
}
