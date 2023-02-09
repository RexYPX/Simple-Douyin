package main

//

import (
	"Simple-Douyin/cmd/user/service"
	"Simple-Douyin/kitex_gen/user"
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

	resp, err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = 2
		resp.StatusMsg = "创建失败"
		return resp, nil
	}

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

	id, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.StatusCode = 2
		resp.StatusMsg = "用户不存在，请注册"
		return resp, nil
	}
	resp = &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserId:     id,
		// 仅在前端获取Token
		// Token: mw.TOKEN,
	}
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp, err = service.NewQueryUserService(ctx).QueryUser(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
