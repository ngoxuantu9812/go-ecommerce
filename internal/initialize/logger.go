package initialize

import (
	"go-ecomm/global"
	"go-ecomm/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
