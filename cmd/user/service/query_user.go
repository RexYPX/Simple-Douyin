package service

import (
	"context"

	"Simple-Douyin/pkg/errno"

	"Simple-Douyin/cmd/user/dal/db"
	"Simple-Douyin/cmd/user/rpc"
	"Simple-Douyin/kitex_gen/relation"
	user "Simple-Douyin/kitex_gen/user"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

// query user info
func (s *QueryUserService) QueryUser(req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	id := req.UserId
	//users, err := db.QueryInfo(s.ctx, id)
	users, err := db.QueryInfoCache(s.ctx, id)
	if err != nil {
		return new(user.UserInfoResponse), err
	}
	if len(users) == 0 {
		return new(user.UserInfoResponse), errno.AuthorizationFailedErr
	}
	u := users[0]

	followCount, err := rpc.FollowCount(s.ctx, &relation.RelationFollowCountRequest{
		UserId: u.Id,
	})
	if err != nil {
		return new(user.UserInfoResponse), err
	}

	followerCount, err := rpc.FollowerCount(s.ctx, &relation.RelationFollowerCountRequest{
		UserId: u.Id,
	})
	if err != nil {
		return new(user.UserInfoResponse), err
	}

	isFollow, err := rpc.IsFollow(s.ctx, &relation.RelationIsFollowRequest{
		UserId:   u.Id,
		ToUserId: req.MUserId,
	})
	if err != nil {
		return new(user.UserInfoResponse), err
	}

	// TODO: finish RPC
	// isFollow, err := rpc.IsFollow(s.ctx, &relation.RelationIsFollowRequest{
	// 	UserId: u.Id,
	// 	ToUserId: ,
	// })
	// if err != nil {
	// 	return new(user.UserInfoResponse), err
	// }

	resp := &user.UserInfoResponse{
		StatusCode:    0,
		Name:          u.Username,
		StatusMsg:     "success",
		Id:            u.Id,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}

	return resp, nil
}
