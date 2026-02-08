package initialize

import (
	"fmt"
	"time"

	"github.com/huyle/go-ecommerce-backend-api/global"
	"github.com/huyle/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func initMySQL() {
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "initMySQL init error")
	global.Logger.Info("initMySQL connected successfully")
	global.Mdb = db

	//set pool
	SetPool()
	//migrate tables
	// migrateTables()
	genTableDAO()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("MySQL error: %s::", err)
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConMaxLifeTime) * time.Second)
}

func genTableDAO() {
	//gen table dao code
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",                                                // specify the path to generate code
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// // Generate the code
	// g.Execute()
}

func migrateTables() {
	er := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if er != nil {
		fmt.Println("migrateTables error:", er)
	}
}
