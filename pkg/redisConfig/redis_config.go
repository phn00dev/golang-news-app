package redisconfig

import (
	"context"
	"fmt"

	"github.com/phn00dev/golang-news-app/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	config *config.Config
}

var RedisClient *redis.Client

func NewRedisConfig(config *config.Config) *RedisConfig {
	return &RedisConfig{
		config: config,
	}
}

func (redisConfig *RedisConfig) Connect() error {
	host := redisConfig.config.RedisConfig.Host
	port := redisConfig.config.RedisConfig.Port
	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Println("Connected to Redis successfully!")
	RedisClient = client
	return nil
}

func (redisConfig *RedisConfig) GetClient() *redis.Client {
	return RedisClient
}
