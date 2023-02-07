package service

import (
	"Simple-Douyin/cmd/feed/pack"
	"Simple-Douyin/cmd/feed/rpc"
	"Simple-Douyin/kitex_gen/feed"
	"context"
	"log"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *feed.FeedRequest) (int64, []*feed.Video, error) {
	log.Println("[ypx debug] service prepare to feed publishVideos")
	publishVideos, err := rpc.GetVideo(req.LatestTime)
	if err != nil {
		log.Println("[ypx debug] service feed publishVideos get error")
		return 0, nil, err
	}
	log.Println("[ypx debug] service feed publishVideos success")
	videos := pack.Videos(publishVideos)

	if len(videos) == 0 {
		log.Println("[ypx debug] len(videos) == 0")
		next_time := time.Now().Unix()
		return next_time, videos, nil
	}

	// TODO: db.Video.CreateTime
	next_time, err := rpc.GetVideoTime(videos[len(videos)-1].Id)
	if err != nil {
		log.Println("[ypx debug] next_time get error")
		return 0, nil, err
	}
	return next_time, videos, nil
}
