package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"memorandum/config"
	"memorandum/pkg/util"
	"strconv"
)

var RDB *redis.Client

func InitRedis() {
	db, err := strconv.Atoi(config.RedisDB)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "",
		DB:       db,
	})
	pong, err := RDB.Ping().Result()
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	fmt.Println("Connected to Redis:", pong)
}
