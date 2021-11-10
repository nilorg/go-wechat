package module

import (
	"github.com/nilorg/go-wechat/v2/proxy/module/config"
	"github.com/nilorg/go-wechat/v2/proxy/module/logger"
	"github.com/nilorg/go-wechat/v2/proxy/module/store"
)

// Init 初始化 module
func Init() {
	config.Init()
	logger.Init()
	store.Init()
}

func Close() {
	logger.Sync()
}
