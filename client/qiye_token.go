package client

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type QiyeTokener interface {
	GetAccessToken() string
}

// QiyeTokenFromRedis 根据Redis客户端创建QiyeToken
type QiyeTokenFromRedis struct {
	opts QiyeTokenClientFromRedisOptions
}

// NewQiyeTokenFromRedis 创建客户端
func NewQiyeTokenFromRedis(opts ...QiyeTokenClientFromRedisOption) *QiyeTokenFromRedis {
	token := &QiyeTokenFromRedis{
		opts: NewQiyeTokenClientFromRedisOptions(opts...),
	}
	return token
}

// GetAccessQiyeToken 获取AccessQiyeToken
func (token *QiyeTokenFromRedis) GetAccessToken() string {
	return token.getRedisValue(token.opts.AccessTokenKey)
}

// getRedisValue 根据Key获取Redis中的字符串
func (client *QiyeTokenFromRedis) getRedisValue(key string) string {
	bytes, err := client.opts.RedisClient.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}

// QiyeTokenForProxy 访问代理不需要使用QiyeToken
type QiyeTokenForProxy struct {
}

// GetAccessQiyeToken 获取AccessToken
func (token *QiyeTokenForProxy) GetAccessToken() string {
	return ""
}
