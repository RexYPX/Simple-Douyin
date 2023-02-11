namespace go publish

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

// 发布视频
struct PublishActionRequest {
    1: i64 user_id
    2: binary data
    3: string title
}

struct PublishActionResponse {
    1: i32 status_code
    2: string status_msg
}

// 发布列表
struct PublishListRequest {
    1: i64 user_id
}

struct PublishListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
}

//video_id[]  to video_list
struct Ids2ListRequest{
    1:list<i64> video_id
    2:i64 user_id
}
struct Ids2ListResponse{
    1:list<Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    Ids2ListResponse PublishIds2List(1:Ids2ListRequest req)
}