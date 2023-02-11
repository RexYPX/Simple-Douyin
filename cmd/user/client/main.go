package main

//

import (
	"Simple-Douyin/pkg/constants"
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"

	"Simple-Douyin/kitex_gen/user"
	"Simple-Douyin/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	client, err := userservice.NewClient(
		constants.UserTableName,
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	//client, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	for {
		//userregister 测试（有mysql）
		//req := &user.UserRegisterRequest{Username: "kirsury", Password: "123456"}
		//resp, err := client.UserRegister(context.Background(), req)

		//userregister 测试（有mysql）
		//req := &user.UserLoginRequest{Username: "kirsury", Password: "123456"}
		//resp, err := client.UserLogin(context.Background(), req)

		//userinfo 测试(无mysql)
		req := &user.UserInfoRequest{UserId: 1}
		resp, err := client.UserInfo(context.Background(), req)

		//
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
