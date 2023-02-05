package rpc

import (
	"Simple-Douyin/kitex_gen/relation"
	"strconv"
)

func Token2Id(token string) (int64, error) {
	userId, err := strconv.Atoi(token)
	if err != nil {
		return 0, err
	}

	return int64(userId), nil
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
