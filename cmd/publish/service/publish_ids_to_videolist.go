package service

import (
	"Simple-Douyin/cmd/publish/dal/db"
	"Simple-Douyin/cmd/publish/rpc"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/favorite"
	publish "Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/kitex_gen/user"
	"context"
)

type PublishIds2ListService struct {
	ctx context.Context
}

// NewPublishIds2ListService new PublishIds2ListService
func NewPublishIds2ListService(ctx context.Context) *PublishIds2ListService {
	return &PublishIds2ListService{ctx: ctx}
}

// video_id[]  to video_list
func (s *PublishIds2ListService) PublishIds2List(req *publish.Ids2ListRequest) (*publish.Ids2ListResponse, error) {
	video_ids := db.PublishIds2List(s.ctx, req.VideoId)

	resp := new(publish.Ids2ListResponse)
	var respVideos []*publish.Video
	for _, v := range video_ids {
		author, _ := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId: v.UserId,
		})

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
			UserId:  int64(req.UserId),
			VideoId: int64(v.ID),
		})

		commentCount, _ := rpc.CommentCount(s.ctx, &comment.CommentListRequest{
			UserId: author.Id,
		})

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

	resp.VideoList = respVideos
	return resp, nil
}
