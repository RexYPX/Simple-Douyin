package main

import (
	"Simple-Douyin/cmd/relation/service"
	relation "Simple-Douyin/kitex_gen/relation"
	"context"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	if req.ToUserId < 0 || req.ActionType <= 0 || req.ActionType > 2 {
		resp.StatusCode = -1
		resp.StatusMsg = "Relation Action request inValid"
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Relation Action fail"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Relation Action success"

	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	if req.UserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Relation Followlist request inValid"
		return resp, nil
	}

	var users []*relation.User
	users, err = service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Relation Followlist fail"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Relation Followlist success"
	resp.UserList = users

	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	if req.UserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Relation Followerlist request inValid"
		return resp, nil
	}

	var users []*relation.User
	users, err = service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "粉丝列表查询失败"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "粉丝列表查询成功"
	resp.UserList = users

	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)

	if req.UserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Relation Friendlist request inValid"
		return resp, nil
	}

	var users []*relation.User
	users, err = service.NewRelationFriendListService(ctx).RelationFriendList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "好友列表查询失败"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "好友列表查询成功"
	resp.UserList = users

	return resp, nil
}
