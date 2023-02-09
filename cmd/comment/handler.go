package main

import (
	"Simple-Douyin/cmd/comment/pack"
	"Simple-Douyin/cmd/comment/service"
	comment "Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/pkg/errno"
	"context"
	"fmt"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentActionResponse)

	if req.UserId < 0 || req.VideoId <= 0 || (req.ActionType <= 0 || req.ActionType >= 3) || (req.ActionType == 1 && len(req.CommentText) == 0) || (req.ActionType == 2 && req.CommentId <= 0) {
		fmt.Println("GQY DEBUG", req.UserId, req.VideoId, req.ActionType, req.CommentText, req.CommentId)
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	comment, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = comment

	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentListResponse)

	if req.UserId < 0 || req.VideoId <= 0 { // warning: Guest users can still get token and access the comments list
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentList = commentList

	return resp, nil
}
