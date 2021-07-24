# 默认使用

# 修改配置文件
```yaml
redis:
  address: "localhost:6379"
  password: ""
  db: 0

apps:
  - id: xxxx # app_id
    secret: "ssss" # app_secret
    refresh_duration: 3600 #（单位秒） 每次间隔刷新时间
    redis_access_token_key: "nilorg:wechat:xxxx:token" #在redis存储access_token的Key
    redis_js_api_ticket_key: "nilorg:wechat:xxxx:ticket" # 在redis存储js_api_ticketKey
```

# Docker使用

[wechat-auto-refresh-server](https://hub.docker.com/r/nilorg/wechat-auto-refresh-server)

```bash
docker run --name wechat-auto-refresh-server \
-e WECHAT_REFRESH_CONFIG="./config.yaml" \
-d nilorg/wechat-auto-refresh-server:latest
```

# k8s使用例子

```bash
kubectl create ns nilorg
# 修改deployment.yaml中的配置文件
kubectl apply -f deployment.yaml
```