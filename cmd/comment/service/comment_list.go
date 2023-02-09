package service

import (
	"Simple-Douyin/cmd/comment/dal/db"
	"Simple-Douyin/cmd/comment/pack"
	"Simple-Douyin/cmd/comment/rpc"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/user"
	"context"
	// "Simple-Douyin/cmd/comment/rpc"
)

type CommentListService struct {
	ctx context.Context
}

// NewQueryNoteService new QueryNoteService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

// QueryNoteService query list of note info
func (s *CommentListService) CommentList(req *comment.CommentListRequest) ([]*comment.Comment, error) {
	commentModels, err := db.QueryComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments := pack.Comments(commentModels)
	for i := 0; i < len(comments); i++ {
		uId := commentModels[i].UserId
		u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{UserId: uId})
		if err != nil {
			return nil, err
		}

		comments[i].User = &comment.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
		}
	}

	return comments, nil
}
