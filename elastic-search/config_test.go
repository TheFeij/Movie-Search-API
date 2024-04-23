package elastic_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadConfig tests loadConfig function
func TestLoadConfig(t *testing.T) {
	require.Equal(t, "http://127.0.0.1:9200", config.ElasticSearchAddress)
}
