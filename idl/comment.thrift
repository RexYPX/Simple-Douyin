namespace go comment

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct BaseResp {
    1: i64 status_code
    2: string status_msg
}

// 评论操作
struct Comment {
  	1: i64 id // 视频评论id
	2: User user // 评论用户信息
   	3: string content // 评论内容
   	4: string create_date // 评论发布日期，格式 mm-dd
}

struct CommentActionRequest {
    1: i64 user_id
	2: i64 video_id
	3: i32 action_type
	4: string comment_text
	5: i64 comment_id
}

struct CommentActionResponse {
	1: BaseResp base_resp
	2: Comment comment
}

// 评论列表
struct CommentListRequest {
    1: i64 user_id
	2: i64 video_id
}

struct CommentListResponse {
	1: BaseResp base_resp
	2: list<Comment> comment_list
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req)
    CommentListResponse CommentList(1: CommentListRequest req)
}