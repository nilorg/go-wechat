# go-wechat
🎉 微信公众号SDK


# Usage
```bash
go get -u github.com/nilorg/go-wechat
```
# Import
```go
import "github.com/nilorg/go-wechat"
```
# Coding
```go
var (
	//client 会自动刷新token
	client wechat.Clienter
)

func init() {
	appID := "wx000000"
	appSecret := "aaaaabbbbbcccc"

	client = wechat.NewClient(appID, appSecret)
}
```

# Context
```go
context := wechat.NewContext(context.Background(), client)
client, err := wechat.FromContext(context)
if err != nil {
	fmt.Println(err)
}
client.GetAccessToken()
client.GetJsAPITicket()
```

## 使用全局唯一AccessToken

Fix [获取 access_token 时 AppSecret 错误，或者 access_token 无效。](https://github.com/nilorg/go-wechat/issues/23)

### 客户端
```go
package main

import (
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/nilorg/go-wechat"
	"github.com/nilorg/sdk/signal"
)

var (
	// Redis 缓存
	Redis *redis.Client
	// client 微信客户端
	client wechat.Clienter
)

func init() {
	initRedis()
}

func initRedis() {
	// 初始化Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:3679",
		Password: "xxxxx",
		DB:       1,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Init redis connection failed: %s \n", err)
		return
	}
	Redis = client
}

func main() {
	client = wechat.NewClientFromRedis()
	// 使用自定义Redis
	// client = wechat.NewClientFromRedis(wechat.ClientFromRedisOptionRedisClient(Redis))
	// 使用自定义RedisKey
	// client = wechat.NewClientFromRedis(
	// 	wechat.ClientFromRedisOptionRedisClient(Redis),
	// 	wechat.ClientFromRedisOptionAccessTokenKey("test_access_token"),
	// 	wechat.ClientFromRedisOptionJsAPITicketKey("test_js_api_ticket"),
	// )
	// 获取内容
	log.Printf("AccessToken:%s\n", client.GetAccessToken())
	log.Printf("JsAPITicket:%s\n", client.GetJsAPITicket())

	signal.AwaitExit()
}
```

# 例子
## 上传文件
```go
filename := "test.jpg"
file, err := os.Open(filename)
if err != nil {
	log.Println(err)
	return
}
defer file.Close()

materialA := material.NewMaterial(client)
result, merr := materialA.UploadTempFile(filename, material.TypeImage, file)
if merr != nil {
	log.Println(merr)
}
log.Printf("%+v\n", result)
```

## 发送客服消息
```go
customService := custom.NewCustom(client)
text := custom.NewTextRequest("o7n1T53CxFZ82ztXqBQKqp_XObEo", "这是客服发送的内容")
err := customService.SendText(text)
if err != nil {
	log.Println(err)
}
```