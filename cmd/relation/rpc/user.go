package rpc

import "Simple-Douyin/cmd/relation/kitex_gen/relation"

func Token2Id(token string) (int64, error) {
	return 1, nil
}

func Id2User(id int64) (*relation.User, error) {
	user := new(relation.User)
	user.Id = id
	user.Name = "YPX"
	user.FollowCount = 0
	user.FollowerCount = 1
	user.IsFollow = false

	return user, nil
}