package main

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/nilorg/go-wechat/v2/proxy/module"
	"github.com/nilorg/go-wechat/v2/proxy/server"
)

func init() {
	// 初始化线程数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	module.Init()
}

func main() {

	// 监控系统信号和创建 Context 现在一步搞定
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// 在收到信号的时候，会自动触发 ctx 的 Done ，这个 stop 是不再捕获注册的信号的意思，算是一种释放资源。
	defer stop()
	go server.HTTP(ctx)
	<-ctx.Done()
}
