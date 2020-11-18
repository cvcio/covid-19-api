package redis

import (
	"errors"
	"fmt"

	rdb "github.com/go-redis/redis/v7"
	"github.com/gomodule/redigo/redis"
)

// Memory service
type Memory struct {
	Pool   *redis.Pool
	Client *rdb.Client
}

// NewCachePool return new pool for redis
func NewCachePool(input string) (*Memory, error) {
	if input == "" {
		return nil, errors.New("redis url cannot be empty")
	}
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", input)
		},
	}

	// Get a connection
	conn := redispool.Get()
	defer conn.Close()
	// Test the connection
	_, err := conn.Do("PING")
	if err != nil {
		return nil, fmt.Errorf("can't connect to the redis database, got error:\n%v", err)
	}

	return &Memory{
		Pool: redispool,
	}, nil
}

// NewLimitsClient return new client for redis
func NewLimitsClient(input string) (*Memory, error) {
	if input == "" {
		return nil, errors.New("redis url cannot be empty")
	}

	option, err := rdb.ParseURL(input)
	if err != nil {
		return nil, err
	}

	client := rdb.NewClient(option)

	return &Memory{
		Client: client,
	}, nil
}
