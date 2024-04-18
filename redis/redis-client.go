// Package redis provides functions to set json data to redis and get json data from it
package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// redisClient redis client
var redisClient *redis.Client

// initializeRedisClient initializes a redis client
func initializeRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        config.RedisAddress,
		DB:          config.RedisDB,
		DialTimeout: config.RedisDialTimeout,
		ReadTimeout: config.RedisReadTimeout,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to initialize redis:\n%v", err)
	}
}
