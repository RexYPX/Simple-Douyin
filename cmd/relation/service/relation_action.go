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
	"context"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationAction
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction create relation between two people
func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	// TODO: user提供token到user_id的接口
	// TEST
	userID := int64(1)
	// TEST END
	// userID := rpc.Token2Id(req.Token)

	toUserId := req.ToUserId

	actionType := req.ActionType

	relationModel := &db.Relation{
		UserId:   userID,
		ToUserId: toUserId,
	}

	// 关注
	if actionType == 1 {
		return db.CreateRelation(s.ctx, []*db.Relation{relationModel})
	}
	// 取消关注
	return db.DeleteRelation(s.ctx, userID, toUserId)
}
