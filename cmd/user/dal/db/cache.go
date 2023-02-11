package db

import (
	"Simple-Douyin/pkg/constants"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr, // redis地址
		Password: "",                  // redis密码，没有则留空
		DB:       0,                   // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func QueryInfoCache(ctx context.Context, userid int64) ([]*User, error) {
	val, err := redisDb.Get(strconv.FormatInt(userid, 10)).Result()
	if err != nil {
		tmp, err := QueryInfo(ctx, userid)
		if err != nil {
			panic(err)
		}
		v, err := json.Marshal(tmp[0])
		if err != nil {
			panic(err)
		}
		err = redisDb.Set(strconv.FormatInt(userid, 10), string(v), 0).Err()
		if err != nil {
			panic(err)
		}
		return tmp, nil
	}
	v := new(User)
	err = json.Unmarshal([]byte(val), &v)
	if err != nil {
		panic(err)
	}
	res := make([]*User, 0)
	res = append(res, v)
	return res, nil
}
