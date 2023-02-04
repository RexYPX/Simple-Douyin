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
	token := test.GetToken()
	addReq1 := &comment.CommentActionRequest{Token: token, VideoId: test.GetVideoId(), ActionType: 1, CommentText: "test1"}
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
	req := &comment.CommentListRequest{Token: test.GetToken(), VideoId: test.GetVideoId()}
	resp, err := c.CommentList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	/* --------------------------------- 删除测试 --------------------------------- */
	// Del Comment1
	// commentId := addResp1.Comment.Id
	token = "1"
	commentId := int64(1)
	delReq := &comment.CommentActionRequest{Token: token, VideoId: test.GetVideoId(), ActionType: 2, CommentId: commentId}
	delResp, err := c.CommentAction(context.Background(), delReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(delResp)
	time.Sleep(time.Second)

	// Get Commentlist
	req = &comment.CommentListRequest{Token: test.GetToken(), VideoId: test.GetVideoId()}
	resp, err = c.CommentList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
