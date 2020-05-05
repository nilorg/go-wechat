package wechat

import (
	"log"

	"github.com/go-redis/redis/v7"
)

// ClientFromRedis Redis客户端
type ClientFromRedis struct {
	opts ClientFromRedisOptions
}

// NewClientFromRedis 创建客户端
func NewClientFromRedis(opts ...ClientFromRedisOption) *ClientFromRedis {
	client := &ClientFromRedis{
		opts: NewClientFromRedisOptions(opts...),
	}
	return client
}

// GetAccessToken 获取AccessToken
func (client *ClientFromRedis) GetAccessToken() string {
	return client.getRedisValue(client.opts.AccessTokenKey)
}

// GetJsAPITicket 获取JsAPITicket
func (client *ClientFromRedis) GetJsAPITicket() string {
	return client.getRedisValue(client.opts.JsAPITicketKey)
}

// getRedisValue 根据Key获取Redis中的字符串
func (client *ClientFromRedis) getRedisValue(key string) string {
	bytes, err := client.opts.RedisClient.Get(key).Bytes()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}
