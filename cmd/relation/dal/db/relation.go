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
	"errors"

	"Simple-Douyin/pkg/constants"

	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	// Id       int64 `grom:"primaryKey;autoIncrement" json:"id"`                       // 关注唯一标识符
	UserId   int64 `gorm:"index:idx_member, priority:1, not null" json:"user_id"`    // 发起关注者ID
	ToUserId int64 `gorm:"index:idx_member, priority:2, not null" json:"to_user_id"` // 被关注者ID
}

type User struct {
	//gorm.Model
	//id       int64
	//username string `json:"user_name"`
	//password string `json:"password"`

	Id            int64  `gorm:"primaryKey;autoIncrement" json:"id"`                //用户唯一标志符号
	Username      string `gorm:"type:varchar(128);not null;index" json:"user_name"` //用户名
	Password      string `gorm:"type:varchar(128);not null" json:"password"`        //用户密码
	FollowCount   int64  `gorm:"not null;default:0" json:"follow_count"`            //关注数
	FollowerCount int64  `gorm:"not null;default:0" json:"follower_count"`          //粉丝数
	FavoriteCount int64  `gorm:"not null;default:0" json:"favorite_count"`          //喜欢数
	TotalFavorite int64  `gorm:"not null;default:0" json:"total_favorite"`          //被赞数
	Avatar        string //用户头像链接Url
	Signature     string //用户个性签名
	Encryption    string //使用的加密手段
	Iter          int    //加密算法迭代次数
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}

// CreateRelation create relation
func CreateRelation(ctx context.Context, r *Relation) error {
	if err := DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", r.UserId, r.ToUserId).First(&Relation{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return DB.WithContext(ctx).Create(r).Error
}

// DeleteRelation delete relation
func DeleteRelation(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(&Relation{}).Error
}

// QueryFollowList query followed users' id
func QueryFollowList(ctx context.Context, userId int64) ([]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.ToUserId)
	}

	return resp, nil
}

// QueryFollowerList query users who followed userId
func QueryFollowerList(ctx context.Context, userId int64) ([]int64, error) {
	var relationFound []*Relation

	var resp []int64

	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		resp = append(resp, r.UserId)
	}

	return resp, nil
}

// QueryFriendList query users who followed userId and userId followed him/her
func QueryFriendList(ctx context.Context, userId int64) ([]int64, error) {
	var resp []int64

	var relationFound []*Relation
	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	for _, r := range relationFound {
		var friendFound *Relation
		if err := DB.WithContext(ctx).Where("to_user_id = ? and user_id = ?", r.UserId, userId).Find(&friendFound).Error; err != nil {
			return resp, err
		}
		if friendFound.ToUserId > 0 {
			resp = append(resp, friendFound.ToUserId)
		}
	}

	return resp, nil
}

// QueryFollowCount return number of users who are followed by userId
func QueryFollowCount(ctx context.Context, userId int64) (int64, error) {
	var relationFound []*Relation

	var resp int64

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	resp = int64(len(relationFound))

	return resp, nil
}

// QueryFollowerCount return number of users who followed userId
func QueryFollowerCount(ctx context.Context, userId int64) (int64, error) {
	var relationFound []*Relation

	var resp int64

	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&relationFound).Error; err != nil {
		return resp, err
	}

	resp = int64(len(relationFound))

	return resp, nil
}

// QueryIsFollow return whether userId followed toUserId
func QueryIsFollow(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	var relationFound *Relation

	var resp bool

	if err := DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", userId, toUserId).First(&relationFound).Error; err != nil {
		resp = false
		return resp, nil
	}

	resp = true

	return resp, nil
}
