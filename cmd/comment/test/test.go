package test

import (
	"Simple-Douyin/kitex_gen/user"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
)

func TokenToUserId(token string) int64 {
	userId, err := strconv.Atoi(token)
	if err != nil {
		return 0
	}

	return int64(userId)
}

func UserIdToToken(userId int64) string {
	token := strconv.FormatInt(userId, 10)

	return token
}

var userIdSequence = int64(1)

func GetUserID() int64 {
	userId := userIdSequence
	atomic.AddInt64(&userIdSequence, 1)

	return userId
}

func GetToken() string {
	return UserIdToToken(GetUserID())
}

func GetVideoId() int64 {
	return 1
}

// 测试接口 GetUser 此接口用于从 user_id 获取一个 user 的信息
func GetUser(ctx context.Context, req *user.UserInfoRequest) (*user.User, error) {
	str := strconv.FormatInt(req.UserId, 10)
	return &user.User{
		Id:            req.UserId,
		Name:          "testuser" + str,
		FollowCount:   req.UserId*2 + 7,
		FollowerCount: req.UserId + 11,
		IsFollow:      false,
	}, nil
}

func PrintGQYDebug() {
	if _, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("\033[1;33;47m%s\033[0m", "[GQY DEBUG]")
		fmt.Printf("file: %s, line %d\n", file, line)
	}
}
