package main

import (
	"Simple-Douyin/kitex_gen/favorite"
	"Simple-Douyin/kitex_gen/favorite/favoriteservice"
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := favoriteservice.NewClient("favorite", client.WithHostPorts("0.0.0.0:8801"))
	if err != nil {
		log.Fatal(err)
	}

	//点赞视频10
	thumbup10req := &favorite.FavoriteActionRequest{Token: "token", VideoId: 10, ActionType: 1}
	thumbup10res, err := c.FavoriteAction(context.Background(), thumbup10req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbup10res)
	time.Sleep(time.Second)

	//拉取点赞视频
	listreq1 := &favorite.FavoriteListRequest{UserId: 1, Token: "token"}
	listres1, err := c.FavoriteList(context.Background(), listreq1, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(listres1)
	time.Sleep(time.Second)

	//取消点赞视频10
	thumbdown10req := &favorite.FavoriteActionRequest{Token: "token", VideoId: 10, ActionType: 2}
	thumbdown10res, err := c.FavoriteAction(context.Background(), thumbdown10req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbdown10res)
	time.Sleep(time.Second)

	//拉取点赞视频
	listreq2 := &favorite.FavoriteListRequest{UserId: 1, Token: "token"}
	listres2, err := c.FavoriteList(context.Background(), listreq2, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(listres2)
	time.Sleep(time.Second)

}
