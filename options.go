package redisqueue

import (
	"time"
)

type Options struct {
	RedisAddr     string
	RedisPassword string
	DB            int
	QueueName     string
	Timeout       time.Duration
}

func NewDefaultOptions() *Options {
	return &Options{
		RedisAddr:     "localhost:6379", // Default Redis address
		RedisPassword: "",               // Default no password
		DB:            0,                // Default database
		QueueName:     "default_queue",  // Default queue name
		Timeout:       5 * time.Second,  // Default timeout
	}
}
func (o *Options) SetRedisAddr(addr string) *Options {
	o.RedisAddr = addr
	return o
}
func (o *Options) SetRedisPassword(password string) *Options {
	o.RedisPassword = password
	return o
}
func (o *Options) SetDB(db int) *Options {
	o.DB = db
	return o
}
func (o *Options) SetQueueName(name string) *Options {
	o.QueueName = name
	return o
}
func (o *Options) SetTimeout(timeout time.Duration) *Options {
	o.Timeout = timeout
	return o
}
func (o *Options) GetRedisAddr() string {
	return o.RedisAddr
}

func (o *Options) GetDB() int {
	return o.DB
}
func (o *Options) GetQueueName() string {
	return o.QueueName
}
func (o *Options) GetTimeout() time.Duration {
	return o.Timeout
}
func (o *Options) Validate() error {
	if o.RedisAddr == "" {
		return ErrInvalidRedisAddr
	}
	if o.QueueName == "" {
		return ErrInvalidQueueName
	}
	if o.Timeout <= 0 {
		return ErrInvalidTimeout
	}
	return nil
}
