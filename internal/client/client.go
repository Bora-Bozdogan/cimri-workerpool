package client

import (
	"context"
	"time"
)

type QueueClient interface {
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	Close() error
}

type WorkerServiceClient struct {
	client QueueClient
	ctx    context.Context
}

func NewWorkerServiceClient(client QueueClient) WorkerServiceClient {
	return WorkerServiceClient{client: client, ctx: context.Background()}
}

func (w WorkerServiceClient) PullItem(queueName string) (string, error) {
	res, err := w.client.BLPop(w.ctx, 5*time.Second, queueName)

	if err == nil && len(res) > 1 {
		return res[1], nil
	}

	return "", err
}
