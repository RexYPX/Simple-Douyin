package main

import (
	user "Simple-Douyin/cmd/user/kitex_gen/user"
	"Simple-Douyin/cmd/user/service"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "用户名或密码为空"
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = 2
		resp.StatusMsg = "用户已存在"
		return resp, nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "用户名或密码为空"
		return resp, nil
	}

	_, err = service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.StatusCode = 2
		resp.StatusMsg = "用户不存在，请注册"
		return resp, nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResponse)
	resp.StatusCode = 1
	resp.StatusMsg = "success"
	resp.Id = 1
	resp.Name = "root"
	resp.FollowCount = 6
	resp.FollowerCount = 6
	resp.IsFollow = true
	return resp, nil
}
