package service

import (
	"Simple-Douyin/cmd/publish/dal/db"
	"Simple-Douyin/cmd/publish/rpc"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/publish"
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
			UserId: v.UserId,
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

		favoriteList, err := rpc.FavoriteList(s.ctx, &favorite.FavoriteListRequest{
			UserId: author.Id,
		})
		if err != nil {
			log.Println("[ypx debug] Gorm rpc.FavoriteList err", err)
			return nil, err
		}

		isFavorite := false
		for _, vid := range favoriteList.VideoList {
			if vid.Id == int64(v.ID) {
				isFavorite = true
				break
			}
		}

		commentCount, err := rpc.CommentCount(s.ctx, &comment.CommentListRequest{
			UserId: author.Id,
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
			FavoriteCount: int64(len(favoriteList.VideoList)),
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		}
		respVideos = append(respVideos, &pv)

	}

	log.Println("[ypx debug] Gorm db.PublishList success")
	return respVideos, nil
}
