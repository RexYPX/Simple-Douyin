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

	// TODO: 添加req的合法性检验

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "关注操作失败"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "关注操作成功"

	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	// TODO: 添加req的合法性检验
	var users []*relation.User
	users, err = service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "关注列表查询失败"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "关注列表查询成功"
	resp.UserList = users

	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	// TODO: 添加req的合法性检验
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

	// TODO: 添加req的合法性检验
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
