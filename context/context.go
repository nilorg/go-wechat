package context

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	wechat "github.com/nilorg/go-wechat"
)

// Context 上下文
type Context struct {
	AppID       string
	AppSecret   string
	accessToken string
	jsAPITicket string
	ticker      *time.Ticker
	rwMutex     *sync.RWMutex // 读写锁
}

// startTicker 启动打点任务
func (c *Context) startTicker() {
	go func() { // 异步
		for range c.ticker.C {
			c.refreshAccessToken() // 刷新AccessToken
			c.refreshJsAPITicket()
		}
	}()
}

// NewContext 创建上下文
func NewContext(appID, appSecret string) *Context {
	context := &Context{
		AppID:     appID,
		AppSecret: appSecret,
		ticker:    time.NewTicker(time.Hour), // 1小时执行一次
		rwMutex:   &sync.RWMutex{},
	}
	context.refreshAccessToken() // 刷新AccessToken
	context.refreshJsAPITicket()
	context.startTicker()
	return context
}

// accessTokenReply ...
type accessTokenReply struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 获取 accessToken
func (c *Context) GetAccessToken() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.accessToken
}

// refreshAccessToken ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func (c *Context) refreshAccessToken() {
	result, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]string{
		"appid":      c.AppID,
		"secret":     c.AppSecret,
		"grant_type": "client_credential",
	})
	if err != nil {
		log.Printf("刷新AccessToken错误：%v", err)
		c.accessToken = ""
	}
	reply := new(accessTokenReply)
	json.Unmarshal(result, reply)

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.accessToken = reply.AccessToken
}

// jsapiTicketReply ...
type jsapiTicketReply struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// GetJsAPITicket 获取 js api ticket
func (c *Context) GetJsAPITicket() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	return c.jsAPITicket
}

// refreshJsAPITicket ...
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
func (c *Context) refreshJsAPITicket() {
	result, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket", map[string]string{
		"access_token": c.GetAccessToken(),
		"type":         "jsapi",
	})
	if err != nil {
		log.Printf("刷新Ticket错误：%v", err)
		c.jsAPITicket = ""
	}
	reply := new(jsapiTicketReply)
	json.Unmarshal(result, reply)

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.jsAPITicket = reply.Ticket
}
