package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wangyupo/GGB/global"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.GGB_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GGB_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
		return
	}

	global.GGB_LOG.Info("redis connect success, here is ping response:", zap.String("pong", pong))
	global.GGB_REDIS = client
}
