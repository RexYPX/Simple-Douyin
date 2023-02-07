package pack

import (
	"Simple-Douyin/cmd/message/dal/db"
	"Simple-Douyin/kitex_gen/message"
	"time"
)

// db.Message -> message.Message
func Message(msg *db.Message) *message.Message {
	if msg == nil {
		return nil
	}

	return &message.Message{
		Id:         int64(msg.MessageId),
		Content:    msg.Content,
		CreateTime: time.Unix(msg.CreateTime, 0).Format("2006-01-02 15:04:05"),
	}
}

// []db.Message -> []message.Message
func Messages(dbmsgs []*db.Message) []*message.Message {
	msgs := make([]*message.Message, 0)

	for _, m := range dbmsgs {
		if msg := Message(m); msg != nil {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}
