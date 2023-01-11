package cache

import (
	"blog/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis(setting *config.App) (*redis.Client,error) {

	client := redis.NewClient(&redis.Options{
		Addr:       setting.Redis.Addr,
		Password:   setting.Redis.Password,
		DB:         setting.Redis.Db,
		MaxRetries: setting.Redis.MaxRetry,
		PoolSize:   setting.Redis.PoolSize,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client,nil
}
