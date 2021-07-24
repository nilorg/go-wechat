package client

import (
	"github.com/go-redis/redis/v8"
)

// TokenClientFromRedisOptions 可选参数列表
type TokenClientFromRedisOptions struct {
	RedisClient    *redis.Client
	AccessTokenKey string
	JsAPITicketKey string
}

// TokenClientFromRedisOption 为可选参数赋值的函数
type TokenClientFromRedisOption func(*TokenClientFromRedisOptions)

const (
	// RedisAccessTokenKey ...
	RedisAccessTokenKey = "github.com/nilorg/go-wechat/access_token"
	// RedisJsAPITicketKey ...
	RedisJsAPITicketKey = "github.com/nilorg/go-wechat/js_api_ticket"
)

// NewTokenClientFromRedisOptions 创建可选参数
func NewTokenClientFromRedisOptions(opts ...TokenClientFromRedisOption) TokenClientFromRedisOptions {
	opt := TokenClientFromRedisOptions{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		AccessTokenKey: RedisAccessTokenKey,
		JsAPITicketKey: RedisJsAPITicketKey,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// ClientFromRedisOptionRedisClient ...
func ClientFromRedisOptionRedisClient(client *redis.Client) TokenClientFromRedisOption {
	return func(o *TokenClientFromRedisOptions) {
		o.RedisClient = client
	}
}

// ClientFromRedisOptionAccessTokenKey ...
func ClientFromRedisOptionAccessTokenKey(accessTokenKey string) TokenClientFromRedisOption {
	return func(o *TokenClientFromRedisOptions) {
		o.AccessTokenKey = accessTokenKey
	}
}

// ClientFromRedisOptionJsAPITicketKey ...
func ClientFromRedisOptionJsAPITicketKey(jsAPITicketKey string) TokenClientFromRedisOption {
	return func(o *TokenClientFromRedisOptions) {
		o.JsAPITicketKey = jsAPITicketKey
	}
}
