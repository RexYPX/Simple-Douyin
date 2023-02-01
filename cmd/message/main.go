package main

import (
	message "Simple-Douyin/cmd/message/kitex_gen/message/meassgeservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MeassgeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
