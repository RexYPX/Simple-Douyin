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
	"Simple-Douyin/cmd/relation/dal/db"
	"Simple-Douyin/cmd/relation/kitex_gen/relation"
	"Simple-Douyin/cmd/relation/rpc"
	"context"
)

type RelationFriendListService struct {
	ctx context.Context
}

// NewRelationFriendListService new RelationFriendList
func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

// RelationFriendList return users who followed each other
func (s *RelationFriendListService) RelationFriendList(req *relation.RelationFriendListRequest) ([]*relation.User, error) {
	var resp []*relation.User

	userID := req.UserId
	// TODO: 使用token验证用户合法性
	token := req.Token
	_, err := rpc.Token2Id(token)
	if err != nil {
		return resp, err
	}

	friendIDs, err := db.QueryFriendList(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	for _, id := range friendIDs {
		// TODO: 使用id获取用户数据
		user, err := rpc.Id2User(id)
		if err != nil {
			return resp, err
		}
		resp = append(resp, user)
	}

	return resp, nil
}
