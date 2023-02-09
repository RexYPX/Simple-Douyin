package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"

	"Simple-Douyin/kitex_gen/user"
	"Simple-Douyin/kitex_gen/user/userservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// RegisterUser create user info
func RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

// LoginUser login user info
func LoginUser(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

// UserInfo query user info
func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
