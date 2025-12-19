package initialize

import (
	"context"
	"flow-blog/pkg/utils"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

//type RedisClient struct{}

func InitRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.DB"),
	})

	_, redisErr := redisCli.Ping(context.Background()).Result()
	if redisErr != nil {
		utils.RecordError("redis连接错误:", redisErr)
	}
	fmt.Println("redis连接成功")
	return redisCli
}
