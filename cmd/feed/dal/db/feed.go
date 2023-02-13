package db

import (
	"Simple-Douyin/pkg/constants"
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId     int64  `json:"user_id"`
	PlayUrl    string `json:"play_url"`
	CoverUrl   string `json:"cover_url"`
	CreateTime int64  `gorm:"default:0"`
	Title      string `json:"title"`
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

// QueryVideoFromTime query get list of video by latest_time
func QueryVideoFromTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var res []*Video
	if latestTime < 0 {
		log.Println("[ypx debug] Gorm QueryVideoFromTime Request err latestTime < 0")
		return res, nil
	}

	newTime := latestTime
	if latestTime == 0 {
		log.Println("[ypx debug] Gorm QueryVideoFromTime Request err latestTime == 0")
		newTime = time.Now().Unix()
	}

	// fix: descending order
	if err := DB.WithContext(ctx).Where("create_time <= ?", newTime).Order("id desc").Limit(constants.MaxFeed).Find(&res).Error; err != nil {
		log.Println("[ypx debug] Gorm QueryVideoFromTime err", err)
		return res, err
	}
	log.Println("[ypx debug] Gorm QueryVideoFromTime success")
	return res, nil
}
