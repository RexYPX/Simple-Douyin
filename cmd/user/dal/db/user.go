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
	"log"

	"Simple-Douyin/pkg/constants"

	"gorm.io/gorm"
)

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

func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Create(users).Error; err != nil {
		return nil, err
	}
	if err := DB.WithContext(ctx).Where("username = ?", users[0].Username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUser query list of user info by name
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryInfo query list of user info by id
func QueryInfo(ctx context.Context, userid int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

type Relation struct {
	gorm.Model
	// Id       int64 `grom:"primaryKey;autoIncrement" json:"id"`                       // 关注唯一标识符
	UserId   int64 `gorm:"index:idx_member, priority:1, not null" json:"user_id"`    // 发起关注者ID
	ToUserId int64 `gorm:"index:idx_member, priority:2, not null" json:"to_user_id"` // 被关注者ID
}

func QueryIsFollow(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	err := DB.Table("relation").Where("user_id = ? and to_user_id = ?", userId, toUserId).First(&Relation{}).Error
	if err != nil {
		log.Println("[ypx debug] user query isFollow err ", err)
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}
