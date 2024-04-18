package api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadConfig tests loadConfig function
func TestLoadConfig(t *testing.T) {
	require.Equal(t, "localhost:8080", config.ServerAddress)
}
