package initialize

import (
	"database/sql"
	"fmt"
	"go-ecomm/global"
	"go-ecomm/internal/model"
	"go.uber.org/zap"
	"gorm.io/gen"
	"time"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(fmt.Sprintf("%s:%d", m.Host, m.Port))
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "Init Mysql initialization error")
	global.Logger.Info("Initialization successfully")
	global.Mdbc = db

	// set Pool
	SetPool()
	migrateTables()

	genTableDAO()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxIdleTime(time.Duration(m.ConnMaxLifetime))
}

func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
		//&po.User{},
		//&po.Role{},
		&model.GoCrmUserV2{},
	)

	if err != nil {
		fmt.Println("migrate tables failed", zap.Error(err))
	}

}

func genTableDAOC() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})

	//g.GenerateModel("go_crm_user_v2")
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
