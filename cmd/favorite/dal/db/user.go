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

	"fmt"
)

type User struct {
	gorm.Model
	// Id int64 `gorm:"primaryKey;autoIncrement" json:id`
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
	VideoId  int64 `json:"video_id"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// // MGetUsers multiple get list of user info
// func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
// 	res := make([]*User, 0)
// 	if len(userIDs) == 0 {
// 		return res, nil
// 	}

// 	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func Delete(ctx context.Context, usrid int64, tousrid int64, video_id int64) error {
	return DB.WithContext(ctx).Where("user_id = ? and to_user_id = ? and video_id = ?", usrid, tousrid, video_id).Delete(&User{}).Error
}

// add favorite
func Add(ctx context.Context, usrid int64, tousrid int64, video_id int64) error {
	res := new(User)

	//如果存在，不再增加
	if err := DB.WithContext(ctx).Where("user_id = ? and to_user_id = ? and video_id = ?", usrid, tousrid, video_id).Find(&res).Error; err != nil {
		fmt.Println("res:", res)
		fmt.Println("db.Add 如果存在，不再增加", video_id)
		return nil
	}

	//如果不存在，增加
	res.UserId = usrid
	res.ToUserId = tousrid
	res.VideoId = video_id
	if err := DB.WithContext(ctx).Create(res).Error; err != nil {
		fmt.Println("db.Add 增加失败", video_id)
		return err
	}
	fmt.Println("db.Add 增加成功")
	return nil
}

// getlist
func QueryUsr(ctx context.Context, usrid int64) ([]int64, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", usrid).Find(&res).Error; err != nil {
		return nil, err
	}

	var video_id []int64
	for _, users := range res {
		video_id = append(video_id, users.VideoId)
	}

	return video_id, nil
}
