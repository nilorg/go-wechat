package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/nilorg/go-wechat"
	"github.com/nilorg/pkg/logger"
	"github.com/nilorg/sdk/convert"
	"github.com/nilorg/sdk/signal"
	"github.com/pkg/errors"
)

const (
	envRedisAddrKey        = "REDIS_ADDR"
	envRedisPasswordKey    = "REDIS_PASSWORD"
	envRedisDbKey          = "REDIS_DB"
	envAppIDKey            = "WECHAT_APP_ID"
	envAppSecretKey        = "WECHAT_APP_SECRET"
	envRefreshDurationKey  = "WECHAT_REFRESH_DURATION"
	envRedisAccessTokenKey = "REDIS_ACCESS_TOKEN_KEY"
	envRedisJsAPITicketKey = "REDIS_JS_API_TICKET_KEY"
)

var (
	redisAddr           string = "127.0.0.1:6379"
	redisPassword       string = ""
	redisDb             int    = 0
	appID               string
	appSecret           string
	refreshDuration     time.Duration = time.Hour
	redisAccessTokenKey               = wechat.RedisAccessTokenKey
	redisJsAPITicketKey               = wechat.RedisJsAPITicketKey
	redisClient         *redis.Client
)

func init() {
	logger.Init()
}

func main() {

	if v := os.Getenv(envRedisAddrKey); v != "" {
		redisAddr = v
	}
	if v := os.Getenv(envRedisPasswordKey); v != "" {
		redisPassword = v
	}
	if v := os.Getenv(envRedisDbKey); v != "" {
		redisDb = convert.ToInt(v)
	}
	if v := os.Getenv(envAppIDKey); v != "" {
		appID = v
	}
	if v := os.Getenv(envAppSecretKey); v != "" {
		appSecret = v
	}
	if v := os.Getenv(envRefreshDurationKey); v != "" {
		refreshDuration = time.Duration(convert.ToInt64(v))
	}
	initRedis()
	logger.Debugln("初始化AccessToken和JsAPITicket")
	refresh()
	ticker := time.NewTicker(refreshDuration)
	go func() { // 异步
		for range ticker.C {
			logger.Debugln("刷新AccessToken和JsAPITicket")
			refresh()
		}
	}()
	signal.AwaitExit()
}

func initRedis() {
	// 初始化Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDb,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logger.Fatalf(
			"Init redis connection failed: %s ",
			errors.Wrap(err, "Ping redis failed"),
		)
	}
	redisClient = client
}

func refresh() {
	token := refreshAccessToken() // 刷新AccessToken
	if token != "" {
		refreshJsAPITicket(token) // 刷新JsAPITicket
	}
}

// refreshAccessToken ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func refreshAccessToken() string {
	result, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]string{
		"appid":      appID,
		"secret":     appSecret,
		"grant_type": "client_credential",
	})
	if err != nil {
		logger.Debugf("刷新AccessToken错误：%v", err)
		return ""
	}
	reply := new(wechat.AccessTokenReply)
	json.Unmarshal(result, reply)

	if err := redisClient.Set(redisAccessTokenKey, reply.AccessToken, 0).Err(); err != nil {
		logger.Errorf("redisClient.Set %s Value: %s Error: %s", redisAccessTokenKey, reply.AccessToken, err)
	}

	return reply.AccessToken
}

// refreshJsAPITicket ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
func refreshJsAPITicket(token string) {
	result, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket", map[string]string{
		"access_token": token,
		"type":         "jsapi",
	})
	if err != nil {
		logger.Debugf("刷新Ticket错误：%v", err)
		return
	}
	reply := new(wechat.JsAPITicketReply)
	json.Unmarshal(result, reply)
	if err := redisClient.Set(redisJsAPITicketKey, reply.Ticket, 0).Err(); err != nil {
		logger.Errorf("redisClient.Set %s Value: %s Error: %s", redisJsAPITicketKey, reply.Ticket, err)
	}
}
