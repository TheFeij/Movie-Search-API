package redis

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestCache tests functions: SetData and GetData
func TestCache(t *testing.T) {
	key := "key"
	value := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key1": "value1",
		},
	}

	t.Run("OK", func(t *testing.T) {
		err := SetData(key, value, 1*time.Minute)
		require.NoError(t, err)

		cachedValue, err := GetData(key)
		require.NoError(t, err)

		require.Equal(t, value, cachedValue)
	})
	t.Run("Expired", func(t *testing.T) {
		err := SetData(key, value, 1*time.Second)
		require.NoError(t, err)

		_, err = GetData(key)
		require.NoError(t, err)

		time.Sleep(1 * time.Second)

		_, err = GetData(key)
		require.Error(t, err)
	})
	t.Run("NonExistent", func(t *testing.T) {
		_, err := GetData("non existent key")
		require.Error(t, err)
	})
}
