package service

import (
	"Simple-Douyin/cmd/publish/dal/db"
	"Simple-Douyin/cmd/publish/rpc"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/favorite"
	publish "Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/kitex_gen/user"
	"context"
	"log"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList publish video.
func (s *PublishListService) PublishList(req *publish.PublishListRequest) (vs []*publish.Video, err error) {
	log.Println("[ypx debug] enter service.PublishList")
	videos, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		log.Println("[ypx debug] kiex service.PublishList err", err)
		return nil, err
	}

	var respVideos []*publish.Video
	for _, v := range videos {
		author, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId:  v.UserId,
			MUserId: req.MUserId,
		})
		if err != nil {
			log.Println("[ypx debug] kitex rpc.GetUser err", err)
			return nil, err
		}

		pAuthor := &publish.User{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      author.IsFollow,
		}

		favorite_count, _ := rpc.FavoriteCount(s.ctx, &favorite.FavoriteCountRequest{
			VideoId: int64(v.ID),
		})

		is_favorite, _ := rpc.IsFavorite(s.ctx, &favorite.IsFavoriteRequest{
			UserId:  int64(req.MUserId),
			VideoId: int64(v.ID),
		})

		commentCount, err := rpc.CommentCount(s.ctx, &comment.CommentListRequest{
			UserId:  author.Id,
			VideoId: int64(v.ID),
		})
		if err != nil {
			log.Println("[ypx debug] Gorm rpc.CommentCount err", err)
			return nil, err
		}

		pv := publish.Video{
			Id:            int64(v.ID),
			Author:        pAuthor,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: favorite_count.FavoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    is_favorite.IsFavorite,
			Title:         v.Title,
		}
		respVideos = append(respVideos, &pv)

	}

	log.Println("[ypx debug] Gorm db.PublishList success")
	return respVideos, nil
}
