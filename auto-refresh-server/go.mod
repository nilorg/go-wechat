module github.com/nilorg/go-wechat/v2/auto-refresh-server

go 1.14

require (
	github.com/go-redis/redis/v8 v8.2.2
	github.com/nilorg/go-wechat/v2 v0.0.0
	github.com/nilorg/pkg v0.0.0-20200923011256-287aac349757
	github.com/nilorg/sdk v0.0.0-20200912025101-a4037e6ee224
	github.com/pkg/errors v0.9.1
)

replace github.com/nilorg/go-wechat/v2 => ../
