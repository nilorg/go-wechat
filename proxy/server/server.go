package server

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nilorg/go-wechat/v2/gateway/middleware"
	"github.com/nilorg/go-wechat/v2/proxy/module/config"
	"github.com/nilorg/go-wechat/v2/proxy/module/store"
	"github.com/sirupsen/logrus"
)

var Transport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func HTTP(ctx context.Context) {
	engine := gin.Default()
	engine.Use(middleware.Header())
	engine.GET("/:appid/*path", checkAppID, proxy)
	engine.POST("/:appid/*path", checkAppID, proxy)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

func proxy(ctx *gin.Context) {
	appID := ctx.Param("appid")
	path := ctx.Param("path")
	logrus.Debugf("APPID:%s,PATH:%s", appID, path)
	appConfig := config.GetApp(appID)
	// 组织要访问的微信接口
	proxyBaseURL := "https://api.weixin.qq.com"
	proxyQuery := ctx.Request.URL.Query()
	switch path {
	case "/sns/oauth2/access_token", "/sns/jscode2session":
		proxyQuery.Set("appid", appConfig.ID)
		proxyQuery.Set("secret", appConfig.Secret)
	case "/sns/oauth2/refresh_token":
		proxyQuery.Set("appid", appConfig.ID)
	case "/cgi-bin/showqrcode":
		proxyBaseURL = "https://mp.weixin.qq.com"
		proxyQuery.Set("ticket", getRedisValue(appConfig.RedisJsAPITicketKey))
	default:
		proxyQuery.Set("access_token", getRedisValue(appConfig.RedisAccessTokenKey))
	}

	var (
		proxyURL *url.URL
		err      error
	)
	proxyURL, err = url.Parse(proxyBaseURL + path)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	logrus.Debugf("请求微信地址：%s?%s", proxyURL, proxyQuery.Encode())
	proxyURL.RawQuery = proxyQuery.Encode()
	proxyReq := *ctx.Request // 复制请求信息
	proxyReq.URL = proxyURL  // 设置代理URL

	//http事务，给一个请求返回一个响应
	var proxyResp *http.Response
	proxyResp, err = Transport.RoundTrip(&proxyReq)
	if err != nil {
		logrus.Errorf("访问源%s错误%v", proxyReq.URL.Host, err)
		ctx.String(http.StatusBadGateway, "请求接口出错")
		return
	}
	defer proxyResp.Body.Close()
	for key, value := range proxyResp.Header { // 设置响应Header
		logrus.Debugf("Header: %s:%v", key, value)
		if strings.EqualFold(key, "Content-Length") || strings.EqualFold(key, "Connection") {
			continue
		}
		for _, v := range value {
			ctx.Writer.Header().Add(key, v)
		}
	}
	ctx.Writer.WriteHeader(proxyResp.StatusCode)
	io.Copy(ctx.Writer, proxyResp.Body)
}

// checkAppID 检查AppID
func checkAppID(ctx *gin.Context) {
	appID := ctx.Param("appid")
	logrus.Debugf("检查AppID:%s是否存在", appID)
	if !config.ExistAppID(appID) {
		logrus.Debugf("未检查到AppID:%s", appID)
		ctx.Status(404)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func getRedisValue(key string) string {
	bytes, err := store.RedisClient.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}
