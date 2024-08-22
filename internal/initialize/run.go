package initialize

import (
	"fmt"
	"go-ecomm/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8082")
}
