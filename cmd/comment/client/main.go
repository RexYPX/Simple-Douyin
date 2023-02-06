package main

import (
	"Simple-Douyin/cmd/comment/test"
	"Simple-Douyin/kitex_gen/comment"
	"Simple-Douyin/kitex_gen/comment/commentservice"
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := commentservice.NewClient("comment", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// CommentActionRequest {Token, VideoId, ActionType, CommentText, CommentId}

	/* --------------------------------- 添加测试 --------------------------------- */
	// Add Comment1
	addReq1 := &comment.CommentActionRequest{Token: test.GetToken(), VideoId: test.GetVideoId(), ActionType: 1, CommentText: "test1"}
	addResp1, err := c.CommentAction(context.Background(), addReq1, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp1)
	time.Sleep(time.Second)

	// // Add Comment2
	addReq2 := &comment.CommentActionRequest{Token: test.GetToken(), VideoId: test.GetVideoId(), ActionType: 1, CommentText: "test2"}
	addResp2, err := c.CommentAction(context.Background(), addReq2, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp2)
	time.Sleep(time.Second)

	// Add Comment3
	addReq3 := &comment.CommentActionRequest{Token: test.GetToken(), VideoId: test.GetVideoId(), ActionType: 1, CommentText: "test3"}
	addResp3, err := c.CommentAction(context.Background(), addReq3, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp3)
	time.Sleep(time.Second)

	// Get Commentlist
	// 其他用户获取列表
	getReq1 := &comment.CommentListRequest{Token: "token", VideoId: test.GetVideoId()}
	getResp1, err := c.CommentList(context.Background(), getReq1, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(getResp1)

	/* --------------------------------- 删除测试 --------------------------------- */
	// Del Comment1
	// commentId := addResp1.Comment.Id
	// 其他用户尝试删除
	delReq1 := &comment.CommentActionRequest{Token: "2", VideoId: test.GetVideoId(), ActionType: 2, CommentId: 1}
	delResp1, err := c.CommentAction(context.Background(), delReq1, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(delResp1)
	time.Sleep(time.Second)

	// 删除本人评论
	delReq2 := &comment.CommentActionRequest{Token: "1", VideoId: test.GetVideoId(), ActionType: 2, CommentId: 1}
	delResp2, err := c.CommentAction(context.Background(), delReq2, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(delResp2)
	time.Sleep(time.Second)

	// Get Commentlist
	getReq2 := &comment.CommentListRequest{Token: "token", VideoId: test.GetVideoId()}
	getResp2, err := c.CommentList(context.Background(), getReq2, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(getResp2)
	time.Sleep(time.Second)

	// token 不合法
	getReq3 := &comment.CommentListRequest{VideoId: test.GetVideoId()}
	getResp3, err := c.CommentList(context.Background(), getReq3, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(getResp3)
}
