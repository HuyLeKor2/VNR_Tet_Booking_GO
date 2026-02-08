package initialize

import (
	"fmt"

	"github.com/huyle/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	//load config
	loadConfig()
	m := global.Config.MySQL
	fmt.Println("Loading config mysql:", m.Username, m.Password)
	initLogger()
	global.Logger.Info("Config Log oke.", zap.String("oke", "success"))
	initMySQL()
	initRedis()

	r := initRouter()
	r.Run(":8082")
}
