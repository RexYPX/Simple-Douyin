// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/publish"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
)

// User pack user info
func User(u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}

	return &userdemo.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}

// publish.Video  to  favorite.Video
func PublishVideo2FavoriteVideo(p_video []*publish.Video) []*favorite.Video {
	f_video := make([]*favorite.Video, 0)

	for _, item := range p_video {

		temp := &favorite.Video{
			Id:            item.Id,
			Author:        (*favorite.User)(item.Author),
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		f_video = append(f_video, temp)
	}

	return f_video
}
