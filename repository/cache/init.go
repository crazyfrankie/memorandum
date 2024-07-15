package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"memorandum/config"
	"memorandum/pkg/util"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "",
		DB:       config.RedisDB,
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	fmt.Println("Connected to Redis:", pong)
}
