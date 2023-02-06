namespace go favorite

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct Video {
    1: i64 id
    2: User author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

// 赞操作
struct FavoriteActionRequest {
    1: string token
	2: i64 video_id
	3: i32 action_type
}

struct FavoriteActionResponse {
    1: i32 status_code
    2: string status_msg
}

// 喜欢列表
struct FavoriteListRequest {
	1: i64 user_id
    2: string token
}

struct FavoriteListResponse {
    1: i32 status_code
    2: string status_msg
	3: list<Video> video_list
}

service FavoriteService {
	FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
	FavoriteListResponse FavoriteList(1: FavoriteListRequest req)
}