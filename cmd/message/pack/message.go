package pack

import (
	"Simple-Douyin/cmd/message/dal/db"
	"Simple-Douyin/kitex_gen/message"
)

// db.Message -> message.Message
func Message(msg *db.Message) *message.Message {
	if msg == nil {
		return nil
	}

	return &message.Message{
		Id:         int64(msg.ID),
		Content:    msg.Content,
		CreateTime: msg.CreateTime,
		FromUserId: msg.UserId,
		ToUserId:   msg.ToUserId,
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
