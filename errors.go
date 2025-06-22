package goredisqueue

import "errors"

var (
	ErrTimeout = errors.New("no message received within the timeout")
)

// ErrInvalidRedisAddr is returned when the Redis address is invalid.
var ErrInvalidRedisAddr = errors.New("invalid Redis address")

// ErrInvalidQueueName is returned when the queue name is invalid.
var ErrInvalidQueueName = errors.New("invalid queue name")

// ErrInvalidTimeout is returned when the timeout is invalid.
var ErrInvalidTimeout = errors.New("invalid timeout")
