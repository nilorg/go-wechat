# 默认使用

```bash
# redis 连接地址
export REDIS_ADDR="127.0.0.1:6379"
# redis 密码
export REDIS_PASSWORD=""
# redis 数据库
export REDIS_DB="0"
# 微信 app id
export WECHAT_APP_ID=""
# 微信 app secret
export WECHAT_APP_SECRET=""
# 微信 刷新 时间(单位：秒)
export WECHAT_REFRESH_DURATION="3600"
# redis 中存储 access_token key
export REDIS_ACCESS_TOKEN_KEY=""
# redis 中存储 js_api_ticket key
export REDIS_JS_API_TICKET_KEY=""

go run main.go
```

# Docker使用

[wechat-auto-refresh-server](https://hub.docker.com/r/nilorg/wechat-auto-refresh-server)

```bash
docker run --name wechat-auto-refresh-server \
-e REDIS_ADDR="127.0.0.1:6379" \
-e WECHAT_APP_ID=xxxx \
-e WECHAT_APP_SECRET=oooo \
-d nilorg/wechat-auto-refresh-server:latest
```