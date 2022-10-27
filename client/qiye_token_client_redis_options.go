package client

import (
	"github.com/go-redis/redis/v8"
)

//  QiyeTokenClientFromRedisOptions 可选参数列表
type QiyeTokenClientFromRedisOptions struct {
	RedisClient    *redis.Client
	AccessTokenKey string
}

//  QiyeTokenClientFromRedisOption 为可选参数赋值的函数
type QiyeTokenClientFromRedisOption func(*QiyeTokenClientFromRedisOptions)

const (
	// QiyeRedisAccessTokenKey ...
	QiyeRedisAccessTokenKey = "github.com/nilorg/go-wechat/qiye/access_token"
)

// NewQiyeTokenClientFromRedisOptions 创建可选参数
func NewQiyeTokenClientFromRedisOptions(opts ...QiyeTokenClientFromRedisOption) QiyeTokenClientFromRedisOptions {
	opt := QiyeTokenClientFromRedisOptions{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		AccessTokenKey: QiyeRedisAccessTokenKey,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// QiyeClientFromRedisOptionRedisClient ...
func QiyeClientFromRedisOptionRedisClient(client *redis.Client) QiyeTokenClientFromRedisOption {
	return func(o *QiyeTokenClientFromRedisOptions) {
		o.RedisClient = client
	}
}

// QiyeClientFromRedisOptionAccessTokenKey ...
func QiyeClientFromRedisOptionAccessTokenKey(accessTokenKey string) QiyeTokenClientFromRedisOption {
	return func(o *QiyeTokenClientFromRedisOptions) {
		o.AccessTokenKey = accessTokenKey
	}
}
