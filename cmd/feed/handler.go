package main

import (
	"Simple-Douyin/cmd/feed/service"
	feed "Simple-Douyin/kitex_gen/feed"
	"context"
	"log"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	log.Println("[ypx debug] kitex handler enter Feed")

	resp = new(feed.FeedResponse)

	feedServiceReq := new(feed.FeedRequest)

	// 用户没有登录 且 第一次打开
	if req.LatestTime < 0 {
		log.Println("[ypx debug] kitex handler req.LatestTime < 0 ")
		feedServiceReq.LatestTime = time.Now().Unix()
	}

	log.Println("[ypx debug] kitex handler prepare to service.NewFeedService(ctx).Feed(feedServiceReq)")
	next_time, videos, err := service.NewFeedService(ctx).Feed(feedServiceReq)
	if err != nil {
		log.Println("[ypx debug] kitex handler service.NewFeedService(ctx).Feed(feedServiceReq) err")
		resp.StatusCode = -1
		resp.StatusMsg = "视频流推送失败"
		return resp, nil
	}
	log.Println("[ypx debug] kitex handler service.NewFeedService(ctx).Feed(feedServiceReq) success")
	resp.StatusCode = 0
	resp.StatusMsg = "视频流推送成功"
	resp.NextTime = next_time
	resp.VideoList = videos

	return resp, nil
}
