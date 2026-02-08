package initialize

import (
	"github.com/huyle/go-ecommerce-backend-api/global"
	"github.com/huyle/go-ecommerce-backend-api/pkg/logger"
)

func initLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
