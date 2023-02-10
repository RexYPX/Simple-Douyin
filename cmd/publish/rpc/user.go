package rpc

import (
	"context"
	"time"

	"Simple-Douyin/kitex_gen/user"
	"Simple-Douyin/kitex_gen/user/userservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"Simple-Douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
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

// GetUser 此接口用于从 user_id 获取一个 user 的信息
func GetUser(ctx context.Context, req *user.UserInfoRequest) (*user.User, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return &user.User{
		Id:            resp.Id,
		Name:          resp.Name,
		FollowCount:   resp.FollowCount,
		FollowerCount: resp.FollowerCount,
		IsFollow:      resp.IsFollow,
	}, nil
}

// func Token2Id(token string) (int64, error) {
// 	userId, err := strconv.Atoi(token)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int64(userId), nil
// }

// func GetFavoriteCount(ctx context.Context, userId int64) (int64, error) {
// 	return 1, nil
// }

// func GetCommentCount(ctx context.Context, userId int64) (int64, error) {
// 	return 1, nil
// }

// func GetIsFavorite(ctx context.Context, userId int64) (bool, error) {
// 	return true, nil
// }
