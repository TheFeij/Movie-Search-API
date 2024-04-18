package rapid_api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadConfig tests loadConfig function
func TestLoadConfig(t *testing.T) {
	require.Equal(t, "24e09e687amsh01656e6e67c47d0p1844bdjsn2c88aeebf80b", config.XRapidAPIKey)
	require.Equal(t, "imdb146.p.rapidapi.com", config.XRapidAPIHost)
}
