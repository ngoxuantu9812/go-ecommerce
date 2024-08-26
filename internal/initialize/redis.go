package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-ecomm/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {

	r := global.Config.Redis

	fmt.Println(fmt.Sprintf("%s:%d", r.Host, r.Port))
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis init fail", zap.Error(err))
		panic(err)
	}

	fmt.Println("Redis init success")
	global.Rds = rdb
	redisExample()

}

func redisExample() {
	err := global.Rds.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		global.Logger.Error("Redis set fail", zap.Error(err))
	}
	value, _ := global.Rds.Get(ctx, "score").Result()
	fmt.Println("Redis score:", value)
}
