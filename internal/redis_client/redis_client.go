package redis_client

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, pass string, num int, protocol int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       num,
		Protocol: protocol,
	})

	return &RedisClient{client: client}
}

// QueueClient interface'ini implement eder
func (r *RedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return r.client.BLPop(ctx, timeout, keys...).Result()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
