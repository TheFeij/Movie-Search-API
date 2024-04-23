package redis

import (
	"context"
	"encoding/json"
	"time"
)

// SetData caches json data (map[string]interface) into redis
func (c Cache) SetData(key string, data map[string]interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = c.redisClient.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetData loads json data (map[string]interface) from cache
func (c Cache) GetData(key string) (map[string]interface{}, error) {
	var result map[string]interface{}

	ctx := context.Background()
	data, err := c.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
