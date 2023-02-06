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
	//拉取
	var userid int64
	//for test
	//req to userid
	userid = 1

	videoids, err := db.QueryUsr(s.ctx, userid)
	if err != nil {
		return nil, err
	}

	//for test
	//rpc : videoids to video_list
	video_list := make([]*favorite.Video, 0)
	if len(videoids) == 1 {

	}

	return video_list, err
}
