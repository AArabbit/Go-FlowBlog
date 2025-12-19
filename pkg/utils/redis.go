package utils

import (
	"context"
	"flow-blog/internal/global"
	"time"
)

// RedisSet 第三个参数传0为永久
func RedisSet(key string, value interface{}, expiration time.Duration) error {
	return global.RedisClient.Set(context.Background(), key, value, expiration).Err()
}

// RedisGet 用key获取value
func RedisGet(key string) (string, error) {
	return global.RedisClient.Get(context.Background(), key).Result()
}

// RedisDel 用key删除数据
func RedisDel(key string) error {
	return global.RedisClient.Del(context.Background(), key).Err()
}
