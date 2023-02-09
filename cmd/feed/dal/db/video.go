package db

import (
	"Simple-Douyin/pkg/constants"
	"context"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	// UpdatedAt time.Time `gorm:"column:update_time;not null;index:idx_update" `
	// Author        User      `gorm:"foreignkey:AuthorID"`
	Author   int64  `gorm:"index:idx_authorid;not null"`
	PlayUrl  string `gorm:"type:varchar(255);not null"`
	CoverUrl string `gorm:"type:varchar(255);not null"`
	// FavoriteCount int    `gorm:"default:0"`
	// CommentCount  int    `gorm:"default:0"`
	Title string `gorm:"type:varchar(50);not null"`
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

// CreateVideo creates a new video
func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
	// err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Create(video).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
	// return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, userId int64) ([]*Video, error) {
	var res []*Video
	if userId <= 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("author = ?", userId).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
	// var pubList []*Video
	// err := DB.WithContext(ctx).Model(&Video{}).Where(&Video{userID: int(userId)}).Find(&pubList).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return pubList, nil
}

func QueryVideo(ctx context.Context, latestTime int64) ([]*Video, error) {
	var res []*Video
	if latestTime <= 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("CreatedAt <= ?", latestTime).Limit(constants.MaxFeed).Order("id desc").Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
