package service

import (
	"Simple-Douyin/cmd/comment/dal/db"
	"Simple-Douyin/pkg/errno"

	// "Simple-Douyin/cmd/comment/rpc"
	"Simple-Douyin/cmd/comment/test"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/user"
	"context"
)

type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentAction comment action
func (s *CommentActionService) CommentAction(req *comment.CommentActionRequest) (*comment.Comment, error) {
	// 发布评论
	if req.ActionType == 1 {
		res := new(comment.Comment)
		// 使用 Token 获取用户信息
		// u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{UserId: test.TokenToUserId(req.Token), Token: req.Token})
		u, err := test.GetUser(s.ctx, &user.UserInfoRequest{UserId: test.TokenToUserId(req.Token), Token: req.Token})
		if err != nil {
			return nil, err
		}

		commentModel := &db.Comment{
			UserId:  u.Id,
			VideoId: req.VideoId,
			Content: req.CommentText,
		}

		err = db.CreateComment(s.ctx, []*db.Comment{commentModel})
		if err != nil {
			return nil, err
		}

		res.Id = int64(commentModel.Model.ID)
		res.User = &comment.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
		}
		res.Content = req.CommentText
		res.CreateDate = commentModel.Model.CreatedAt.Format("01-02")

		return res, nil
	}

	// 删除评论
	// fix: The server needs to check whether there is a delete permission
	commentModel, err := db.GetComment(s.ctx, req.CommentId)
	if err != nil {
		return nil, err
	}

	// 如果删除的发起用户不是该评论的生产用户，不允许发起删除操作
	if commentModel.UserId != test.TokenToUserId(req.Token) {
		return nil, errno.AuthorizationFailedErr
	}

	return nil, db.DeleteComment(s.ctx, req.CommentId)
}
