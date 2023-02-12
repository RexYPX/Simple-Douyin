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
	"Simple-Douyin/cmd/relation/rpc"
	"Simple-Douyin/kitex_gen/relation"
	"Simple-Douyin/kitex_gen/user"
	"context"
)

type RelationFollowListService struct {
	ctx context.Context
}

// NewRelationActionService new RelationAction
func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

// RelationAction create relation between two people
func (s *RelationFollowListService) RelationFollowList(req *relation.RelationFollowListRequest) ([]*relation.User, error) {
	var resp []*relation.User

	userID := req.UserId

	followIDs, err := db.QueryFollowList(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	for _, id := range followIDs {
		u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
			UserId:  id,
			MUserId: req.MUserId,
		})
		if err != nil {
			return resp, err
		}
		resp = append(resp, &relation.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
		})
	}

	return resp, nil
}
