package service

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"

	"Simple-Douyin/cmd/user/dal/db"
	user "Simple-Douyin/cmd/user/kitex_gen/user"
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
	users, err := db.QueryInfo(s.ctx, id)
	if err != nil {
		return new(user.UserInfoResponse), err
	}
	if len(users) == 0 {
		return new(user.UserInfoResponse), errno.AuthorizationFailedErr
	}
	u := users[0]

	resp := &user.UserInfoResponse{
		StatusCode:    0,
		Name:          u.Username,
		StatusMsg:     "success",
		Id:            u.Id,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      true,
	}

	return resp, nil
}
