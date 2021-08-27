package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	set "turan.com/web_demo/settings"
)

var Client *redis.Client

func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", set.RedisC.Host, set.RedisC.Port),
		DB:       set.RedisC.Dbname,
		PoolSize: set.RedisC.PoolSize,
	})
	_, err = Client.Ping().Result()
	if err != nil {
		panic(err)
		return
	}
	return
}

func Close() {
	_ = Client.Close()
}
