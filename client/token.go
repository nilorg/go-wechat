package client

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Tokener interface {
	GetAccessToken() string
	GetJsAPITicket() string
}

// TokenFromRedis 根据Redis客户端创建Token
type TokenFromRedis struct {
	opts TokenClientFromRedisOptions
}

// NewTokenFromRedis 创建客户端
func NewTokenFromRedis(opts ...TokenClientFromRedisOption) *TokenFromRedis {
	token := &TokenFromRedis{
		opts: NewTokenClientFromRedisOptions(opts...),
	}
	return token
}

// GetAccessToken 获取AccessToken
func (token *TokenFromRedis) GetAccessToken() string {
	return token.getRedisValue(token.opts.AccessTokenKey)
}

// GetJsAPITicket 获取JsAPITicket
func (token *TokenFromRedis) GetJsAPITicket() string {
	return token.getRedisValue(token.opts.JsAPITicketKey)
}

// getRedisValue 根据Key获取Redis中的字符串
func (client *TokenFromRedis) getRedisValue(key string) string {
	bytes, err := client.opts.RedisClient.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}

// TokenForProxy 访问代理不需要使用Token
type TokenForProxy struct {
}

// GetAccessToken 获取AccessToken
func (token *TokenForProxy) GetAccessToken() string {
	return ""
}

// GetJsAPITicket 获取JsAPITicket
func (token *TokenForProxy) GetJsAPITicket() string {
	return ""
}
