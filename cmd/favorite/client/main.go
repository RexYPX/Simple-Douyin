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
	c, err := favoriteservice.NewClient("favorite", client.WithHostPorts("0.0.0.0:8989"))
	if err != nil {
		log.Fatal(err)
	}

	//5点赞视频5
	thumbup10req := &favorite.FavoriteActionRequest{UserId: 5, VideoId: 5, ActionType: 1}
	thumbup10res, err := c.FavoriteAction(context.Background(), thumbup10req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbup10res)
	time.Sleep(time.Second)

	//5点赞视频50
	thumbup20req := &favorite.FavoriteActionRequest{UserId: 5, VideoId: 50, ActionType: 1}
	thumbup20res, err := c.FavoriteAction(context.Background(), thumbup20req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbup20res)
	time.Sleep(time.Second)

	//5点赞视频500
	thumbup200req := &favorite.FavoriteActionRequest{UserId: 5, VideoId: 500, ActionType: 1}
	thumbup200res, err := c.FavoriteAction(context.Background(), thumbup200req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbup200res)
	time.Sleep(time.Second)

	//5拉取点赞视频
	listreq1 := &favorite.FavoriteListRequest{UserId: 5}
	listres1, err := c.FavoriteList(context.Background(), listreq1, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(listres1)
	time.Sleep(time.Second)

	//5取消点赞视频5
	thumbdown10req := &favorite.FavoriteActionRequest{UserId: 5, VideoId: 5, ActionType: 2}
	thumbdown10res, err := c.FavoriteAction(context.Background(), thumbdown10req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbdown10res)
	time.Sleep(time.Second)

	//5拉取点赞视频
	listreq2 := &favorite.FavoriteListRequest{UserId: 5}
	listres2, err := c.FavoriteList(context.Background(), listreq2, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(listres2)
	time.Sleep(time.Second)

	//50点赞视频500
	thumbup50req := &favorite.FavoriteActionRequest{UserId: 50, VideoId: 500, ActionType: 1}
	thumbup50res, err := c.FavoriteAction(context.Background(), thumbup50req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(thumbup50res)
	time.Sleep(time.Second)

	//get FavoriteCount   视频500
	favorite_countreq500 := &favorite.FavoriteCountRequest{VideoId: 500}
	favorite_countres500, err := c.FavoriteCount(context.Background(), favorite_countreq500, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("favorite_countres500  shoud 2 :", favorite_countres500)
	time.Sleep(time.Second)

	//get FavoriteCount   视频50
	favorite_countreq50 := &favorite.FavoriteCountRequest{VideoId: 50}
	favorite_countres50, err := c.FavoriteCount(context.Background(), favorite_countreq50, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("favorite_countres50  shoud 1 :", favorite_countres50)
	time.Sleep(time.Second)

	// userid like video_id or not
	//5 do not like 5
	fivenotlinke5req := &favorite.IsFavoriteRequest{UserId: 5, VideoId: 5}
	fivenotlinke5res, err := c.IsFavorite(context.Background(), fivenotlinke5req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("5 do not like 5 :", fivenotlinke5res)

	//5   like 50
	fivelinke50req := &favorite.IsFavoriteRequest{UserId: 5, VideoId: 50}
	fivelinke50res, err := c.IsFavorite(context.Background(), fivelinke50req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("5   like 50 :", fivelinke50res)
	time.Sleep(time.Second)

}
