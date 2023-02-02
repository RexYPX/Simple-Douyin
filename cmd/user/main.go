package main

import (
	user "Simple-Douyin/cmd/user/kitex_gen/user/userservice"
	"log"

	"Simple-Douyin/cmd/user/dal"
)

func main() {
	dal.Init()

	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
