package cache

import (
	"JT_CLUB/conf"
	"fmt"
	"github.com/go-redis/redis"
)

var Cache *redis.Client

func InitCache() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Config.Redis.Host, conf.Config.Redis.Port),
		Password: conf.Config.Redis.Password, // no password set
		DB:       conf.Config.Redis.DB,       // use default DB
	})
	_, err := Cache.Ping().Result()
	if err != nil {
		panic(err)
	}
}
