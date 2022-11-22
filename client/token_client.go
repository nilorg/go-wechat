package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// TokenClient 客户端
type TokenClient struct {
	BaseURL     string
	appID       string
	appSecret   string
	accessToken string
	jsAPITicket string
	Proxy       bool
	ticker      *time.Ticker
	rwMutex     *sync.RWMutex // 读写锁
	HttpClient  *http.Client
}

// startTicker 启动打点任务
func (c *TokenClient) startTicker() {
	go func() { // 异步
		for range c.ticker.C {
			log.Println("刷新AccessToken和JsAPITicket")
			c.refreshAccessToken() // 刷新AccessToken
			c.refreshJsAPITicket()
		}
	}()
}

// NewTokenClient 创建客户端
func NewTokenClient(appID, appSecret string) *TokenClient {
	client := &TokenClient{
		BaseURL:    "https://api.weixin.qq.com",
		Proxy:      false,
		appID:      appID,
		appSecret:  appSecret,
		ticker:     time.NewTicker(time.Hour), // 1小时执行一次
		rwMutex:    &sync.RWMutex{},
		HttpClient: http.DefaultClient,
	}
	client.refreshAccessToken() // 刷新AccessToken
	client.refreshJsAPITicket()
	client.startTicker()
	return client
}

// AccessTokenReply ...
type AccessTokenReply struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 获取 accessToken
func (c *TokenClient) GetAccessToken() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.accessToken
}

// refreshAccessToken ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func (c *TokenClient) refreshAccessToken() {
	url := fmt.Sprintf("%s/cgi-bin/token", c.BaseURL)
	value := map[string]string{
		"grant_type": "client_credential",
	}
	if !c.Proxy {
		value["appid"] = c.appID
		value["secret"] = c.appSecret
	}
	result, err := Get(c.HttpClient, url, value)
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

// JsAPITicketReply ...
type JsAPITicketReply struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// GetJsAPITicket 获取 js api ticket
func (c *TokenClient) GetJsAPITicket() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.jsAPITicket
}

// refreshJsAPITicket ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
func (c *TokenClient) refreshJsAPITicket() {
	url := fmt.Sprintf("%s/cgi-bin/ticket/getticket", c.BaseURL)
	value := map[string]string{
		"type": "jsapi",
	}
	if !c.Proxy {
		value["access_token"] = c.GetAccessToken()
	}
	result, err := Get(c.HttpClient, url, value)
	if err != nil {
		log.Printf("刷新Ticket错误：%v", err)
		// c.jsAPITicket = ""
		return
	}
	reply := new(JsAPITicketReply)
	json.Unmarshal(result, reply)

	c.rwMutex.Lock()
	c.jsAPITicket = reply.Ticket
	c.rwMutex.Unlock()
}
