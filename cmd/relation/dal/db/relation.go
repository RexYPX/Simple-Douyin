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

package db

import (
	"context"

	"Simple-Douyin/pkg/constants"

	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	Id       int64 `grom:"primaryKey;autoIncrement" json:"id"`                       // 关注唯一标识符
	UserId   int64 `gorm:"index:idx_member, priority:1, not null" json:"user_id"`    // 发起关注者ID
	ToUserId int64 `gorm:"index:idx_member, priority:2, not null" json:"to_user_id"` // 被关注者ID
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}

// CreateRelation create relation
func CreateRelation(ctx context.Context, relations []*Relation) error {
	return DB.WithContext(ctx).Create(relations).Error
}

// DeleteRelation delete relation
func DeleteRelation(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(&Relation{}).Error
}

// QueryFollowList query followed users' id
func QueryFollowList(ctx context.Context, userId int64) (*[]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return &resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.ToUserId)
	}

	return &resp, nil
}

// QueryFollowerList query users who followed userId
func QueryFollowerList(ctx context.Context, userId int64) (*[]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return &resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.ToUserId)
	}

	return &resp, nil
}

// QueryFriendList query users who followed userId and userId followed him/her
func QueryFriendList(ctx context.Context, userId int64) (*[]int64, error) {
	var resp []int64

	var relationFound []*Relation
	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return &resp, err
	}

	var friendFound []*Relation
	for _, r := range relationFound {
		if err := DB.WithContext(ctx).Where("to_user_id = ? and user_id = ?", userId, r.UserId).Find(&friendFound).Error; err != nil {
			return &resp, err
		}
	}

	for _, r := range friendFound {
		resp = append(resp, r.UserId)
	}

	return &resp, nil
}
