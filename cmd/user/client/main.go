package main

//

import (
	"context"
	"log"
	"time"

	"Simple-Douyin/cmd/user/kitex_gen/user"
	"Simple-Douyin/cmd/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		//userregister 测试（有mysql）
		//req := &user.UserRegisterRequest{Username: "kirsury", Password: "123456"}
		//resp, err := client.UserRegister(context.Background(), req)

		//userregister 测试（有mysql）
		req := &user.UserLoginRequest{Username: "kirsury", Password: "123456"}
		resp, err := client.UserLogin(context.Background(), req)

		//userinfo 测试(无mysql)
		//req := &user.UserInfoRequest{UserId: "1", Token: "2"}
		//resp, err := client.UserInfo(context.Background(), req)

		//
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
