package mredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	Pool *redis.Pool
}

var (
	pool             *redis.Pool
	RedisCacheClient *RedisCache
)

func init() {
	pool = newPool("127.0.0.1:6379", "")
	RedisCacheClient = &RedisCache{Pool: pool}
}

func newPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		MaxActive:   1024,
		IdleTimeout: 600 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialPassword(password), redis.DialReadTimeout(5*time.Second), redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
