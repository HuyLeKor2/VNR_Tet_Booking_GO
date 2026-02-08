package initialize

import (
	"context"
	"fmt"

	"github.com/huyle/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func initRedis() {
	r := global.Config.Redis
	// refer https://redis.uptrace.dev/guide/go-redis.html for details
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host + fmt.Sprintf(":%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("initRedis init error", zap.Error(err))
		panic(err)
	}

	fmt.Println("initRedis connected successfully")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", "100", 0).Err()
	if err != nil {
		fmt.Println("RedisExample Set error:", err)
		return
	}
	val, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("RedisExample Get error:", err)
		return
	}
	global.Logger.Info("RedisExample Get score:", zap.String("score", val))
}
