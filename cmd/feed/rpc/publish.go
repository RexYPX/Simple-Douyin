package rpc

import (
	"time"

	"Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/kitex_gen/publish/publishservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
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
	publishClient = c
}

// GetVideo 此接口用于从 create_at 获取一个 video 的信息
func GetVideo(int64) ([]*publish.Video, error) {
	author := publish.User{
		Id:            1,
		Name:          "YPX",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	return []*publish.Video{
		{
			Id:            1,
			Author:        &author,
			PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         "DemoVideo",
		},
	}, nil
}

// GetVideoTime 此接口用于从 video_id 获取 对应 video 的创建时间
func GetVideoTime(int64) (int64, error) {
	return time.Now().Unix(), nil
}
