package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// InitializeRedisClient initializes a redis client
func InitializeRedisClient() Cacher {
	cache := Cache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:        config.RedisAddress,
			DB:          config.RedisDB,
			DialTimeout: config.RedisDialTimeout,
			ReadTimeout: config.RedisReadTimeout,
		}),
	}

	ctx := context.Background()
	if err := cache.redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to initialize redis:\n%v", err)
	}

	return cache
}
