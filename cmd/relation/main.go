package main

import (
	relation "Simple-Douyin/cmd/relation/kitex_gen/relation/relationservice"
	"log"

	"Simple-Douyin/cmd/relation/dal"
)

func Init() {
	dal.Init()
}

func main() {

	// r, err := etcd.NewEtcdRegistry([]string{constants.ETCDAddress})
	// if err != nil {
	// 	panic(err)
	// }

	// addr, err := net.ResolveTCPAddr(constants.TCP, constants.RelationServiceAddr)
	// if err != nil {
	// 	panic(err)
	// }

	Init()

	// svr := relation.NewServer(new(RelationServiceImpl),
	// 	server.WithRegistry(r),
	// 	server.WithServiceAddr(addr),
	// )

	// err = svr.Run()

	svr := relation.NewServer(new(RelationServiceImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
