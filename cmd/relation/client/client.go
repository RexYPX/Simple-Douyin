package main

import (
	"context"
	"log"
	"time"

	"Simple-Douyin/cmd/relation/kitex_gen/relation"
	"Simple-Douyin/cmd/relation/kitex_gen/relation/relationservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := relationservice.NewClient("relation", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	toUserID := int64(5)
	for i := 0; i < 3; i++ {
		//RelationAction 测试（有mysql）
		// 关注
		// req := &relation.RelationActionRequest{Token: "YPX", ToUserId: toUserID, ActionType: 1}
		// resp, err := client.RelationAction(context.Background(), req)
		// 取关
		// req := &relation.RelationActionRequest{Token: "YPX", ToUserId: toUserID, ActionType: 0}
		// resp, err := client.RelationAction(context.Background(), req)

		//RelationFollowList 测试（有mysql）
		// req := &relation.RelationFollowListRequest{UserId: 1, Token: "YPX"}
		// resp, err := client.RelationFollowList(context.Background(), req)

		//RelationFollowerList 测试（有mysql）
		// req := &relation.RelationFollowerListRequest{UserId: 1, Token: "YPX"}
		// resp, err := client.RelationFollowerList(context.Background(), req)

		//RelationFriendList 测试（有mysql）
		req := &relation.RelationFriendListRequest{UserId: 1, Token: "YPX"}
		resp, err := client.RelationFriendList(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)

		toUserID++
		time.Sleep(time.Second)
	}
}
