package redis

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestLoadConfig tests loadConfig function
func TestLoadConfig(t *testing.T) {
	require.Equal(t, "127.0.0.1:6379", config.RedisAddress)
	require.Equal(t, 0, config.RedisDB)
	require.Equal(t, 100*time.Millisecond, config.RedisDialTimeout)
	require.Equal(t, 100*time.Millisecond, config.RedisReadTimeout)
}
