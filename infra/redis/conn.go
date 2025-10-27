package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb  *redis.Client
	once sync.Once
	ctx  context.Context
)

// NewRedisServer initializes and returns a Redis client
func NewRedisServer() *redis.Client {
	once.Do(func() {
		ctx = context.Background()

		// Initialize Redis client
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6380", // Redis server address
			Password: "",               // no password set
			DB:       0,                // default DB
		})

		// Test connection
		if pong, err := rdb.Ping(ctx).Result(); err != nil {
			panic(fmt.Sprintf("Failed to connect Redis: %v", err))
		} else {
			fmt.Println("Connected to Redis:", pong)
		}
	})

	return rdb
}
