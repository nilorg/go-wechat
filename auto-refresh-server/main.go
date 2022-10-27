package main

import (
	"context"
	"encoding/json"
	"os/signal"
	"syscall"
	"time"

	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module"
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/config"
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/logger"
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/store"
	"github.com/nilorg/go-wechat/v2/client"
)

func init() {
	module.Init()
}

func main() {
	defer module.Close()
	// 监控系统信号和创建 Context 现在一步搞定
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// 在收到信号的时候，会自动触发 ctx 的 Done ，这个 stop 是不再捕获注册的信号的意思，算是一种释放资源。
	defer stop()

	apps := config.GetApps()
	tickers := make([]*time.Ticker, len(apps))
	defer func() {
		for _, t := range tickers {
			t.Stop()
		}
	}()
	for i, app := range apps {
		logger.Sugared.Debugf("初始化App:[%s]的AccessToken和JsAPITicket", app.ID)
		refresh(app.ID, app.Secret, app.RedisAccessTokenKey, app.RedisJsAPITicketKey, app.Type)
		time.Sleep(time.Second)
		ticker := time.NewTicker(time.Duration(app.RefreshDuration) * time.Second)
		go func(a *config.AppConfig) { // 异步
			for range ticker.C {
				logger.Sugared.Debugf("App:[%s]刷新AccessToken和JsAPITicket", a.ID)
				refresh(a.ID, a.Secret, a.RedisAccessTokenKey, a.RedisJsAPITicketKey, app.Type)
			}
		}(app)
		tickers[i] = ticker
	}
	<-ctx.Done()
}

func refresh(appID string, appSecret string, redisAccessTokenKey, redisJsAPITicketKey []string, typ string) {
	if typ == "qiye" {
		refreshQiyeAccessToken(appID, appSecret, redisAccessTokenKey) // 刷新AccessToken
	} else {
		token := refreshAccessToken(appID, appSecret, redisAccessTokenKey) // 刷新AccessToken
		if token != "" {
			refreshJsAPITicket(appID, token, redisJsAPITicketKey) // 刷新JsAPITicket
		}
	}
}

// refreshAccessToken ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func refreshAccessToken(appID, appSecret string, redisAccessTokenKey []string) string {
	result, err := client.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]string{
		"appid":      appID,
		"secret":     appSecret,
		"grant_type": "client_credential",
	})
	if err != nil {
		logger.Sugared.Debugf("App:[%s]刷新AccessToken错误：%v", appID, err)
		return ""
	}
	reply := new(client.AccessTokenReply)
	json.Unmarshal(result, reply)
	for _, v := range redisAccessTokenKey {
		if err := store.RedisClient.Set(context.Background(), v, reply.AccessToken, time.Second*time.Duration(reply.ExpiresIn)).Err(); err != nil {
			logger.Sugared.Error("App:[%s]redisClient.Set %s Value: %s Error: %s", appID, v, reply.AccessToken, err)
		}
	}
	logger.Sugared.Debugf("App:[%s]最新AccessToken: %s", appID, reply.AccessToken)
	return reply.AccessToken
}

// refreshJsAPITicket ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
func refreshJsAPITicket(appID, token string, redisJsAPITicketKey []string) {
	result, err := client.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket", map[string]string{
		"access_token": token,
		"type":         "jsapi",
	})
	if err != nil {
		logger.Sugared.Debugf("App:[%s]刷新Ticket错误：%v", appID, err)
		return
	}
	reply := new(client.JsAPITicketReply)
	json.Unmarshal(result, reply)
	logger.Sugared.Debugf("App:[%s]最新JsAPITicket: %s", appID, reply.Ticket)
	for _, v := range redisJsAPITicketKey {
		if err := store.RedisClient.Set(context.Background(), v, reply.Ticket, time.Second*time.Duration(reply.ExpiresIn)).Err(); err != nil {
			logger.Sugared.Errorf("App:[%s]redisClient.Set %s Value: %s Error: %s", appID, v, reply.Ticket, err)
		}
	}
}

// refreshQiyeAccessToken 刷新企业微信AccessToken
// https://developer.work.weixin.qq.com/document/path/91039
func refreshQiyeAccessToken(appID, appSecret string, redisAccessTokenKey []string) {
	result, err := client.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken", map[string]string{
		"corpid":     appID,
		"corpsecret": appSecret,
	})
	if err != nil {
		logger.Sugared.Debugf("App:企业微信[%s]刷新AccessToken错误：%v", appID, err)
		return
	}
	reply := new(client.AccessTokenReply)
	json.Unmarshal(result, reply)
	for _, v := range redisAccessTokenKey {
		if err := store.RedisClient.Set(context.Background(), v, reply.AccessToken, time.Second*time.Duration(reply.ExpiresIn)).Err(); err != nil {
			logger.Sugared.Error("App:企业微信[%s]redisClient.Set %s Value: %s Error: %s", appID, v, reply.AccessToken, err)
		}
	}
	logger.Sugared.Debugf("App:企业微信[%s]最新AccessToken: %s", appID, reply.AccessToken)
}
