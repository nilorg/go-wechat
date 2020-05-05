# 默认使用

```bash
# redis 连接地址
export REDIS_ADDR="127.0.0.1"
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