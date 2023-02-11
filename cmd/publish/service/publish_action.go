package service

import (
	"Simple-Douyin/cmd/publish/dal/db"
	"Simple-Douyin/cmd/publish/pack"
	"Simple-Douyin/kitex_gen/publish"
	"Simple-Douyin/pkg/constants"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"time"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// CommentAction comment action
func (s *PublishActionService) PublishAction(req *publish.PublishActionRequest) error {
	log.Println("[ypx debug] enter service.PublishAction")
	// 发布视频
	formPos := strings.LastIndex(req.Title, ".")
	title := req.Title[:formPos]

	pwd, err := os.Getwd()
	log.Println("[ypx bebug] ", pwd)
	finalName := fmt.Sprintf("%s/public/video/%d_%s", pwd, req.UserId, req.Title)

	if err != nil {
		log.Println("[ypx debug] kitex PublishAction pwd err ", err)
		return err
	}
	f, err := os.Create(finalName)
	if err != nil {
		log.Println("[ypx debug] kitex PublishAction os.Create err ", err)
		return err
	}
	w := bufio.NewWriter(f)
	wlen, err := w.Write(req.Data)
	if err != nil {
		log.Println("[ypx debug] kitex PublishAction w.Write(req.Data) err ", err)
		return err
	}
	log.Println("write data len: ", wlen)
	defer f.Close()

	coverName := fmt.Sprintf("%s/public/cover/%d_%s", pwd, req.UserId, title)
	// 获取视频封面
	pack.GetSnapshot(finalName, coverName, 1)

	// 获取本机IP
	// ip, err := pack.GetOutBoundIP()
	// if err != nil {
	// 	log.Println("[ypx debug] kitex PublishAction pack.GetOutBoundIP() err ", err)
	// 	return err
	// }

	vName := fmt.Sprintf("%d_%s", req.UserId, req.Title)
	cName := fmt.Sprintf("%d_%s", req.UserId, title)

	playURL := "http://" + constants.HertzServiceIP + constants.FileServerAddr + "/video/" + vName
	coverURL := "http://" + constants.HertzServiceIP + constants.FileServerAddr + "/cover/" + cName + ".png"

	videoModel := &db.Video{
		UserId:     req.UserId,
		PlayUrl:    playURL,
		CoverUrl:   coverURL,
		CreateTime: time.Now().Unix(),
		Title:      title,
	}

	err = db.CreateVideo(s.ctx, []*db.Video{videoModel})
	if err != nil {
		log.Println("[ypx debug] kitex PublishAction db.CreateVideo(s.ctx, []*db.Video{videoModel}) err ", err)
		return err
	}

	log.Println("[ypx debug] kitex PublishAction success")
	return nil
}
