package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/nilorg/go-wechat/v2/gateway/server"
	"github.com/spf13/viper"
)

func init() {
	// 初始化线程数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	configFilename := "./config.yaml"
	if v := os.Getenv("WECHAT_GATEWAY_CONFIG"); v != "" {
		configFilename = v
	}
	viper.SetConfigFile(configFilename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("OnConfigChange: %+v", in)
	})
}

func main() {

	// 监控系统信号和创建 Context 现在一步搞定
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// 在收到信号的时候，会自动触发 ctx 的 Done ，这个 stop 是不再捕获注册的信号的意思，算是一种释放资源。
	defer stop()
	server.HTTP()
	<-ctx.Done()
}
