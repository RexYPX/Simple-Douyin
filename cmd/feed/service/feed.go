package service

import (
	"Simple-Douyin/cmd/feed/dal/db"
	"Simple-Douyin/cmd/feed/rpc"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/feed"
	"Simple-Douyin/kitex_gen/user"
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

	// TODO: 根据 Token 进行个性化feed

	dbVideos, err := db.QueryVideo(s.ctx, req.LatestTime)
	if err != nil {
		log.Println("[ypx debug] service feed dbVideos get error")
		return 0, nil, err
	}
	log.Println("[ypx debug] service feed dbVideos success")

	var videos []*feed.Video
	for _, dbV := range dbVideos {
		user, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId: req.UserId,
		})
		if err != nil {
			return 0, nil, err
		}
		author := feed.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}

		// favoriteList, err := rpc.FavoriteList(s.ctx, &favorite.FavoriteListRequest{
		// 	UserId: author.Id,
		// })
		// if err != nil {
		// 	log.Println("[ypx debug] Gorm rpc.FavoriteList err", err)
		// 	return nil, err
		// }
		favorite_count, _ := rpc.FavoriteCount(s.ctx, &favorite.FavoriteCountRequest{
			VideoId: int64(dbV.ID),
		})

		// isFavorite := false
		// for _, vid := range favoriteList.VideoList {
		// 	if vid.Id == int64(v.ID) {
		// 		isFavorite = true
		// 		break
		// 	}
		// }

		is_favorite, _ := rpc.IsFavorite(s.ctx, &favorite.IsFavoriteRequest{
			UserId:  int64(req.UserId),
			VideoId: int64(dbV.ID),
		})

		commentCount, err := rpc.CommentCount(s.ctx, &comment.CommentListRequest{
			UserId: user.Id,
		})
		if err != nil {
			return 0, nil, err
		}

		v := feed.Video{
			Id:       int64(dbV.ID),
			Author:   &author,
			PlayUrl:  dbV.PlayUrl,
			CoverUrl: dbV.CoverUrl,
			// FavoriteCount: int64(len(favoriteList.VideoList)),
			FavoriteCount: favorite_count.FavoriteCount,
			CommentCount:  commentCount,
			// IsFavorite:    isFavorite,
			IsFavorite: is_favorite.IsFavorite,
			Title:      dbV.Title,
		}
		videos = append(videos, &v)

	}

	if len(videos) == 0 {
		log.Println("[ypx debug] len(videos) == 0")
		next_time := time.Now().Unix()
		return next_time, videos, nil
	}

	next_time := dbVideos[0].CreatedAt.Unix()

	return next_time, videos, nil
}
