package client

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type WorkerServiceClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewWorkerServiceClient(addr string, pass string, num int, protocol int) WorkerServiceClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,     // No password set
		DB:       num,      // Use default DB
		Protocol: protocol, // Connection protocol
	})
	return WorkerServiceClient{client: client, ctx: context.Background()}
}

func (w WorkerServiceClient) PullItem(queueName string) (string, error) {
	res, err := w.client.BLPop(w.ctx, 5*time.Second, queueName).Result()

	if err != nil && len(res) > 1 {
		return res[1], nil
	}

	return "", err
}
