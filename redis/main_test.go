package redis

import (
	"os"
	"testing"
)

// TestMain runs before other tests, initializes redis client.
func TestMain(m *testing.M) {
	initializeRedisClient()

	os.Exit(m.Run())
}
