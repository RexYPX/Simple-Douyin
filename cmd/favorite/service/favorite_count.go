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
	favorite "Simple-Douyin/kitex_gen/favorite"
)

type FavoriteCountService struct {
	ctx context.Context
}

// NewFavoriteCountService new FavoriteCountService
func NewFavoriteCountService(ctx context.Context) *FavoriteCountService {
	return &FavoriteCountService{
		ctx: ctx,
	}
}

func (s *FavoriteCountService) FavoriteCount(req *favorite.FavoriteCountRequest) (int64, error) {
	//拉取
	return db.QueryFavoriteCount(s.ctx, req.VideoId)
}
