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
	"log"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"

	"Simple-Douyin/cmd/favorite/dal/db"
	"Simple-Douyin/cmd/favorite/pack"
	"Simple-Douyin/cmd/favorite/rpc"
	favorite "Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/publish"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// getlist
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	//本用户id + video_id[]  获取 video_list
	video_ids, _ := db.QueryUsr(s.ctx, req.UserId)

	if len(video_ids) == 0 {
		log.Println("FavoriteList : video_ids is blank")
	}

	resp, err := rpc.PublishIds2List(s.ctx, &publish.Ids2ListRequest{VideoId: video_ids, UserId: req.UserId})

	f_video := pack.PublishVideo2FavoriteVideo(resp.VideoList)

	return f_video, err
}
