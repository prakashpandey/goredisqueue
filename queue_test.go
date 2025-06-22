package redisqueue

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestPayload struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

func (tp *TestPayload) Marshal() ([]byte, error) {
	return json.Marshal(tp)
}

func (tp *TestPayload) Unmarshal(data []byte) error {
	return json.Unmarshal(data, tp)
}

func setupQueue(t *testing.T) *Queue {
	t.Helper()
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	return &Queue{
		client:  client,
		queue:   "test_queue",
		timeout: 1,
	}
}

func TestEnqueueDequeue(t *testing.T) {
	ctx := context.Background()
	queue := setupQueue(t)

	send := &TestPayload{ID: "123", Body: "hello"}
	require.NoError(t, queue.Enqueue(ctx, send))

	recv := &TestPayload{}
	err := queue.Dequeue(ctx, recv)
	require.NoError(t, err)

	assert.Equal(t, send.ID, recv.ID)
	assert.Equal(t, send.Body, recv.Body)
}

func TestDequeueTimeout(t *testing.T) {
	ctx := context.Background()
	queue := setupQueue(t)

	// Clean queue before test
	_ = queue.Enqueue(ctx, &TestPayload{ID: "temp", Body: "flush"})
	temp := &TestPayload{}
	_ = queue.Dequeue(ctx, temp)

	recv := &TestPayload{}
	err := queue.Dequeue(ctx, recv)

	assert.ErrorIs(t, err, ErrTimeout)
}
