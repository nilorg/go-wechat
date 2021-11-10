package module

import (
	"github.com/nilorg/go-wechat/v2/gateway/module/config"
	"github.com/nilorg/go-wechat/v2/gateway/module/logger"
	"github.com/nilorg/go-wechat/v2/gateway/module/store"
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
