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
	fmt.Println("获取上下文中的微信客户端成功")
}
client.GetAccessToken()
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