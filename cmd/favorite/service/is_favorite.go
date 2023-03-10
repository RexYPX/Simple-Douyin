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

type IsFavoriteService struct {
	ctx context.Context
}

// NewIsFavoriteService new IsFavoriteService
func NewIsFavoriteService(ctx context.Context) *IsFavoriteService {
	return &IsFavoriteService{
		ctx: ctx,
	}
}

// ueser_id like video_id
func (s *IsFavoriteService) IsFavorite(req *favorite.IsFavoriteRequest) (bool, error) {
	return db.QueryIsFavorite(s.ctx, req)
}
