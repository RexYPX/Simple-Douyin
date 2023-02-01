namespace go feed

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

// 视频相关
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

// 视频流接口
struct FeedRequest {
    1: i64 latest_time
    2: string token
}

struct FeedResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
    4: i64 next_time
}

service FeedService {
    FeedResponse Feed(1: FeedRequest req)
}