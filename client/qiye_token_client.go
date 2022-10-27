package client

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// QiyeTokenClient 客户端
type QiyeTokenClient struct {
	BaseURL     string
	appID       string
	appSecret   string
	accessToken string
	Proxy       bool
	ticker      *time.Ticker
	rwMutex     *sync.RWMutex // 读写锁
}

// startTicker 启动打点任务
func (c *QiyeTokenClient) startTicker() {
	go func() { // 异步
		for range c.ticker.C {
			log.Println("刷新AccessToken和JsAPITicket")
			c.refreshAccessToken() // 刷新AccessToken
		}
	}()
}

// NewQiyeTokenClient 创建客户端
func NewQiyeTokenClient(appID, appSecret string) *QiyeTokenClient {
	client := &QiyeTokenClient{
		BaseURL:   "https://qyapi.weixin.qq.com",
		Proxy:     false,
		appID:     appID,
		appSecret: appSecret,
		ticker:    time.NewTicker(time.Hour), // 1小时执行一次
		rwMutex:   &sync.RWMutex{},
	}
	client.refreshAccessToken() // 刷新AccessToken
	client.startTicker()
	return client
}

// GetAccessToken 获取 accessToken
func (c *QiyeTokenClient) GetAccessToken() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.accessToken
}

// refreshAccessToken ...
// https://developer.work.weixin.qq.com/document/path/91039
func (c *QiyeTokenClient) refreshAccessToken() {
	url := fmt.Sprintf("%s/cgi-bin/gettoken", c.BaseURL)
	value := make(map[string]string)
	if !c.Proxy {
		value["corpid"] = c.appID
		value["corpsecret"] = c.appSecret
	}
	result, err := Get(url, value)
	if err != nil {
		log.Printf("刷新AccessToken错误：%v\n", err)
		// c.accessToken = ""
		return
	}
	reply := new(AccessTokenReply)
	json.Unmarshal(result, reply)
	c.rwMutex.Lock()
	c.accessToken = reply.AccessToken
	c.rwMutex.Unlock()
}
