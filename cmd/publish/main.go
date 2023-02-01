package main

import (
	publish "Simple-Douyin/cmd/publish/kitex_gen/publish/publishservice"
	"log"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
