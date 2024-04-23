package redis

import "time"

// Cacher an interface for caching data.
type Cacher interface {
	// SetData sets the data in the cache with the provided key and expiration duration.
	// returns an error if the operation fails.
	SetData(key string, data map[string]interface{}, expiration time.Duration) error

	// GetData retrieves the data from the cache using the provided key.
	// returns the data associated with the key and an error if the operation fails.
	GetData(key string) (map[string]interface{}, error)
}
