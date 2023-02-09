package pack

import (
	"Simple-Douyin/cmd/feed/dal/db"
	"Simple-Douyin/kitex_gen/feed"
)

// db.Video -> feed.Video
func Video(publishVideo *db.Video) *feed.Video {
	if publishVideo == nil {
		return nil
	}
	pAuthor := publishVideo.Author
	author := &feed.User{
		Id:            pAuthor.Id,
		Name:          pAuthor.Name,
		FollowCount:   pAuthor.FollowCount,
		FollowerCount: pAuthor.FollowerCount,
		IsFollow:      pAuthor.IsFollow,
	}

	return &feed.Video{
		Id:            publishVideo.Id,
		Author:        author,
		PlayUrl:       publishVideo.PlayUrl,
		CoverUrl:      publishVideo.CoverUrl,
		FavoriteCount: publishVideo.FavoriteCount,
		CommentCount:  publishVideo.CommentCount,
		IsFavorite:    publishVideo.IsFavorite,
		Title:         publishVideo.Title,
	}
}

// []db.Video -> []feed.Video
func Videos(publishVideos []*db.Video) []*feed.Video {
	videos := make([]*feed.Video, 0)

	for _, pv := range publishVideos {
		if v := Video(pv); v != nil {
			videos = append(videos, v)
		}
	}

	return videos
}
