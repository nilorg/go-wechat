# go-wechat
🎉 微信公众号SDK

## 使用全局唯一AccessToken

Fix [获取 access_token 时 AppSecret 错误，或者 access_token 无效。](https://github.com/nilorg/go-wechat/issues/23)

### 服务器端

使用[wechat-auto-refresh-server](https://github.com/nilorg/go-wechat/tree/master/auto-refresh-server)作为自动刷新服务器


# 第二版构思

wechat-auto-refresh-server:作为自动刷新Token服务器

client:实现微信接口

gateway:负责解析微信消息推送，分发给各个业务系统。统一入口

proxy:作为clinet的代理，目的主要是为了解决多个（公众号、小程序）配置混乱问题。统一出口