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

type Comment struct {
	gorm.Model
	UserId  int64  `json:"user_id"`
	VideoId int64  `json:"video_id"`
	Content string `json:"content"`
}

func (n *Comment) TableName() string {
	return constants.CommentTableName
}

// CreateComment create comment info
func CreateComment(ctx context.Context, comments []*Comment) error {
	if err := DB.WithContext(ctx).Create(comments).Error; err != nil {
		return err
	}
	return nil
}

// QueryComment query get list of comment info
func QueryComment(ctx context.Context, videoID int64) ([]*Comment, error) {
	var res []*Comment
	if videoID <= 0 {
		return res, nil
	}

	// fix: descending order
	if err := DB.WithContext(ctx).Where("video_id = ?", videoID).Order("id desc").Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// DeleteComment delete comment info
func DeleteComment(ctx context.Context, commentID int64) error {
	return DB.WithContext(ctx).Where("id = ?", commentID).Delete(&Comment{}).Error
}

// GetComment simple get comment info
func GetComment(ctx context.Context, commentID int64) (*Comment, error) {
	var res *Comment
	if commentID <= 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id = ?", commentID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
