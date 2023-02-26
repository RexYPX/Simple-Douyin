namespace go message

// 发送消息
struct MessageActionRequest {
	1: i64 user_id
	2: i64 to_user_id
	3: i32 action_type
	4: string content
}

struct MessageActionResponse {
    1: i32 status_code
    2: string status_msg
}

// 聊天记录
struct Message {
	1: i64 id
    2: string content
	3: i64 create_time
    4: i64 from_user_id
    5: i64 to_user_id
}

struct MessageChatRequest {
	1: i64 user_id
	2: i64 to_user_id
    3: i64 pre_msg_time
}

struct MessageChatResponse {
    1: i32 status_code
    2: string status_msg
	3: list<Message> message_list
}

service MeassgeService {
	MessageActionResponse MessageAction(1: MessageActionRequest req)
	MessageChatResponse MessageChat(1: MessageChatRequest req)
}
