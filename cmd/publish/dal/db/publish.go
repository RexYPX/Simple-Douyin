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
	"log"

	"Simple-Douyin/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId   int64  `json:"user_id"`
	PlayUrl  string `json:"play_url"`
	CoverUrl string `json:"cover_url"`
	Title    string `json:"title"`
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		log.Println("[ypx debug] Gorm CreateVideo err", err)
		return err
	}

	// for _, v := range videos {
	// 	log.Println(v.ID, v.CreatedAt, v.UserId, v.PlayUrl, v.CoverUrl, v.Title)
	// 	if err := DB.WithContext(ctx).Create(v).Error; err != nil {
	// 		log.Println("[ypx debug] Gorm CreateVideo err", err)
	// 		return err
	// 	}
	// }
	return nil
}

// QueryVideo query get list of video info
func QueryVideo(ctx context.Context, videoID int64) ([]*Video, error) {
	var res []*Video
	if videoID <= 0 {
		log.Println("[ypx debug] Gorm QueryVideo err")
		return res, nil
	}

	// fix: descending order
	if err := DB.WithContext(ctx).Where("video_id = ?", videoID).Order("id desc").Find(&res).Error; err != nil {
		log.Println("[ypx debug] Gorm QueryVideo err", err)
		return res, err
	}

	log.Println("[ypx debug] Gorm QueryVideo success")
	return res, nil
}

// PublishList returns a list of videos with userId.
func PublishList(ctx context.Context, userId int64) ([]*Video, error) {
	log.Println("[ypx debug] Gorm enter PublishList")
	var res []*Video
	if userId <= 0 {
		log.Println("[ypx debug] Gorm PublishList Request err")
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Order("id desc").Find(&res).Error; err != nil {
		log.Println("[ypx debug] Gorm PublishList err", err)
		return res, err
	}
	log.Println("[ypx debug] Gorm PublishVideo success", res)
	return res, nil
}

// QueryVideoFromTime query get list of video by latest_time
func QueryVideoFromTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var res []*Video
	if latestTime < 0 {
		log.Println("[ypx debug] Gorm QueryVideoFromTime Request err")
		return res, nil
	}

	// fix: descending order
	if err := DB.WithContext(ctx).Where("CreateAt <= ?", latestTime).Order("id desc").Limit(constants.MaxFeed).Find(&res).Error; err != nil {
		log.Println("[ypx debug] Gorm QueryVideoFromTime err", err)
		return res, err
	}
	log.Println("[ypx debug] Gorm QueryVideoFromTime success")
	return res, nil
}

// videoids  []int64   to db.video  []*Video
func PublishIds2List(ctx context.Context, video_ids []int64) []*Video {
	var res []*Video

	for _, vid := range video_ids {
		temp := make([]*Video, 0)
		DB.WithContext(ctx).Where("id = ?", vid).Find(&temp)
		if len(temp) != 0 {
			res = append(res, temp[0])
		}
	}
	if len(res) == 0 {
		log.Println("[]*Video res is blank!!!!")
	}
	return res
}

// // DeleteComment delete comment info
// func DeleteVideo(ctx context.Context, videoID int64) error {
// 	return DB.WithContext(ctx).Where("id = ?", videoID).Delete(&Video{}).Error
// }

// // GetVideo simple get video info
// func GetVideo(ctx context.Context, videoID int64) (*Video, error) {
// 	var res *Video
// 	if videoID <= 0 {
// 		return res, nil
// 	}

// 	if err := DB.WithContext(ctx).Where("id = ?", videoID).Find(&res).Error; err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }
