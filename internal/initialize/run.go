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
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafka()
	r := InitRouter()
	r.Run(":8002")
}
