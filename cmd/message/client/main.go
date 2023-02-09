package main

import (
	// TODO meassge -> message
	"Simple-Douyin/kitex_gen/message"
	"Simple-Douyin/kitex_gen/message/meassgeservice"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/util"
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := meassgeservice.NewClient(
		constants.MessageServiceName,
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	rand.Seed(time.Now().Unix())

	// send message
	// TODO token
	for i := 0; i < 10; i++ {
		sendMsgReq := &message.MessageActionRequest{
			// Token:      util.GetRandomString(16),
			ToUserId:   int64(util.GetRandomUserID()),
			ActionType: 1,
			Content:    util.GetRandomString(32),
		}
		sendMsgResp, err := c.MessageAction(context.Background(), sendMsgReq)
		if err != nil {
			panic(err)
		}
		log.Println(sendMsgResp)
		time.Sleep(time.Microsecond * 10)
	}

	// get history
	getMsgHistoryReq := &message.MessageChatRequest{
		// Token:    util.GetRandomString(16),
		ToUserId: int64(util.GetRandomUserID()),
	}
	getMsgHistoryResp, err := c.MessageChat(context.Background(), getMsgHistoryReq)
	if err != nil {
		panic(err)
	}
	log.Println(getMsgHistoryResp)
}
