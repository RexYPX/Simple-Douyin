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

package service

import (
	"context"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"

	"Simple-Douyin/cmd/favorite/dal/db"
	favorite "Simple-Douyin/cmd/favorite/kitex_gen/favorite"
)

type ActionFavoriteService struct {
	ctx context.Context
}

// NewActionFavoriteService new ActionFavoriteService
func NewActionFavoriteService(ctx context.Context) *ActionFavoriteService {
	return &ActionFavoriteService{
		ctx: ctx,
	}
}

// like or dislike
func (s *ActionFavoriteService) ActionFavorite(req *favorite.FavoriteActionRequest) error {
	//for test
	var usrid int64
	var tousrid int64

	usrid = 1
	tousrid = 2

	video_id := req.VideoId
	action_type := req.ActionType

	//1点赞
	if action_type == 1 {
		return db.Add(s.ctx, usrid, tousrid, video_id)
	} else {
		//2取消点赞
		return db.Delete(s.ctx, usrid, tousrid, video_id)
	}

}
