package pack

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/message"
)

// db.Message -> message.Message
func Message(msg *message.Message) *api.Message {
	if msg == nil {
		return nil
	}

	return &api.Message{
		ID:         msg.Id,
		Content:    msg.Content,
		CreateTime: msg.CreateTime,
		FromUserID: msg.FromUserId,
		ToUserID:   msg.ToUserId,
	}
}

// []db.Message -> []message.Message
func Messages(dbmsgs []*message.Message) []*api.Message {
	msgs := make([]*api.Message, 0)

	for _, m := range dbmsgs {
		if msg := Message(m); msg != nil {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}
