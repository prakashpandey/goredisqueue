package redisqueue

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Queue struct {
	client  *redis.Client
	queue   string
	timeout time.Duration
}

func New(client *redis.Client, queueName string, timeout time.Duration) *Queue {
	return &Queue{
		client:  client,
		queue:   queueName,
		timeout: timeout,
	}
}

func NewWithOptions(opt Options) *Queue {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.RedisAddr,
		Password: opt.RedisPassword,
		DB:       opt.DB,
	})
	return &Queue{
		client:  client,
		queue:   opt.QueueName,
		timeout: opt.Timeout,
	}
}

func (q *Queue) Enqueue(ctx context.Context, p Payload) error {
	data, err := p.Marshal()
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	return q.client.LPush(ctx, q.queue, data).Err()
}

func (q *Queue) Dequeue(ctx context.Context, p Payload) error {
	result, err := q.client.BRPop(ctx, q.timeout, q.queue).Result()
	if err != nil {
		if err == redis.Nil {
			return ErrTimeout
		}
		return fmt.Errorf("dequeue error: %w", err)
	}
	if len(result) != 2 {
		return fmt.Errorf("unexpected BRPOP result: %v", result)
	}
	return p.Unmarshal([]byte(result[1]))
}
