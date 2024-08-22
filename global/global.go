package global

import (
	"go-ecomm/pkg/logger"
	"go-ecomm/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
