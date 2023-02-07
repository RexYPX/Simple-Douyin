package pack

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/feed"
)

// feed.Video -> api.Video
func Video(feedVideo *feed.Video) *api.Video {
	if feedVideo == nil {
		return nil
	}
	fAuthor := feedVideo.Author
	author := &api.User{
		ID:            fAuthor.Id,
		Name:          fAuthor.Name,
		FollowCount:   fAuthor.FollowCount,
		FollowerCount: fAuthor.FollowerCount,
		IsFollow:      fAuthor.IsFollow,
	}

	return &api.Video{
		ID:            feedVideo.Id,
		Author:        author,
		PlayURL:       feedVideo.PlayUrl,
		CoverURL:      feedVideo.CoverUrl,
		FavoriteCount: feedVideo.FavoriteCount,
		CommentCount:  feedVideo.CommentCount,
		IsFavorite:    feedVideo.IsFavorite,
		Title:         feedVideo.Title,
	}
}

// []feed.Video -> []api.Video
func Videos(feedVideos []*feed.Video) []*api.Video {
	videos := make([]*api.Video, 0)

	for _, fv := range feedVideos {
		if v := Video(fv); v != nil {
			videos = append(videos, v)
		}
	}

	return videos
}
