package module

import (
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/config"
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/logger"
	"github.com/nilorg/go-wechat/v2/auto-refresh-server/module/store"
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
