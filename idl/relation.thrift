namespace go relation

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

// 关注操作
struct RelationActionRequest {
    1: string token
	2: i64 to_user_id
	3: i32 action_type
}

struct RelationActionResponse {
    1: i32 status_code
    2: string status_msg
}

// 关注列表
struct RelationFollowListRequest {
	1: i64 user_id
    2: string token
}

struct RelationFollowListResponse {
    1: i32 status_code
    2: string status_msg
	3: list<User> user_list
}

// 粉丝列表
struct RelationFollowerListRequest {
	1: i64 user_id
    2: string token
}

struct RelationFollowerListResponse {
    1: i32 status_code
    2: string status_msg
	3: list<User> user_list
}

// 好友列表
struct RelationFriendListRequest {
	1: i64 user_id
    2: string token
}

struct RelationFriendListResponse {
    1: i32 status_code
    2: string status_msg
	3: list<User> user_list
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req)
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req)
	RelationFriendListResponse RelationFriendList(1:RelationFriendListRequest req)
}