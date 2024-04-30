package config

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestLoadConfig tests loadConfig function
func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./", "config_test")
	require.NoError(t, err)

	require.Equal(t, "api", config.ServerAddress)
	require.Equal(t, time.Millisecond*100, config.RedisReadTimeout)
	require.Equal(t, time.Millisecond*100, config.RedisDialTimeout)
	require.Equal(t, "key", config.XRapidAPIKey)
	require.Equal(t, "host", config.XRapidAPIHost)
	require.Equal(t, "es8", config.ElasticSearchAddress)
	require.Equal(t, "endpoint", config.RapidAPISearchEndpoint)
	require.Equal(t, "redis", config.RedisAddress)
	require.Equal(t, 0, config.RedisDB)
}
