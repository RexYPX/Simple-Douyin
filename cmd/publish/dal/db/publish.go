package db

import (
	"context"
	"log"
	"time"

	"Simple-Douyin/pkg/constants"

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

// CreateVideo create video info
func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		log.Println("[ypx debug] Gorm CreateVideo err", err)
		return err
	}
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
		log.Println("[ypx debug] Gorm QueryVideoFromTime Request err latestTime < 0")
		return res, nil
	}

	newTime := latestTime
	if latestTime == 0 {
		log.Println("[ypx debug] Gorm QueryVideoFromTime Request err latestTime == 0")
		newTime = time.Now().Unix()
	}

	log.Println("[ypx debug] DB ", DB)

	// fix: descending order
	if err := DB.WithContext(ctx).Where("create_time <= ?", newTime).Order("id desc").Limit(constants.MaxFeed).Find(&res).Error; err != nil {
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
