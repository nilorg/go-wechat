package server

import (
	"context"
	"fmt"
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
	"github.com/spf13/viper"
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

	fmt.Println(appID)
	var (
		proxyURL *url.URL
		err      error
	)
	proxyURL, err = url.Parse(fmt.Sprintf("https://api.weixin.qq.com/%s", path))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	rdsAccessTokenKey := viper.GetString(fmt.Sprintf("apps.%s.redis_access_token_key", appID))
	proxyQuery := proxyURL.Query()
	proxyQuery.Set("access_token", getRedisValue(rdsAccessTokenKey))
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
	if !config.ExistAppID(appID) {
		ctx.Status(400)
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
