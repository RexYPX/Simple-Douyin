package pack

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/publish"
)

// publish.Video -> api.Video
func publishVideo2ApiVideo(publishVideo *publish.Video) *api.Video {
	if publishVideo == nil {
		return nil
	}
	fAuthor := publishVideo.Author
	author := &api.User{
		ID:            fAuthor.Id,
		Name:          fAuthor.Name,
		FollowCount:   fAuthor.FollowCount,
		FollowerCount: fAuthor.FollowerCount,
		IsFollow:      fAuthor.IsFollow,
	}

	return &api.Video{
		ID:            publishVideo.Id,
		Author:        author,
		PlayURL:       publishVideo.PlayUrl,
		CoverURL:      publishVideo.CoverUrl,
		FavoriteCount: publishVideo.FavoriteCount,
		CommentCount:  publishVideo.CommentCount,
		IsFavorite:    publishVideo.IsFavorite,
		Title:         publishVideo.Title,
	}
}

// []publish.Video -> []api.Video
func publishVideos2ApiVideos(publishVideos []*publish.Video) []*api.Video {
	videos := make([]*api.Video, 0)

	for _, pv := range publishVideos {
		if v := publishVideo2ApiVideo(pv); v != nil {
			videos = append(videos, v)
		}
	}

	return videos
}
