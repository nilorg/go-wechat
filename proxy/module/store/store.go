package store

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/nilorg/go-wechat/v2/proxy/module/logger"

	"github.com/spf13/viper"
)

var (
	// RedisClient redis 客户端
	RedisClient *redis.Client
)

// Init 初始化
func Init() {
	initRedis()
}

func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	logger.Standard.Info(pong)
}
