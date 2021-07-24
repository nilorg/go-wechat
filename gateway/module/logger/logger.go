package logger

import (
	"go.uber.org/zap"
)

var (
	Standard *zap.Logger
	Sugared  *zap.SugaredLogger
)

func Init() {
	Standard, _ = zap.NewDevelopment()
	Sugared = Standard.Sugar()
}

func Sync() {
	Sugared.Sync()
	Standard.Sync()
}
