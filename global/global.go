package global

import (
	"github.com/huyle/go-ecommerce-backend-api/pkg/logger"
	"github.com/huyle/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZAP
	Mdb    *gorm.DB
	MySQL  *setting.MySQLSetting
	Rdb    *redis.Client
)

/*
Config
Redis
MySQL
...
*/
