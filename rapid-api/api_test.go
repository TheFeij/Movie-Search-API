package rapid_api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSearchQuery tests SearchQuery
func TestSearchQuery(t *testing.T) {
	result, err := SearchQuery("brad")
	require.NoError(t, err)

	require.NotEmpty(t, result)
}
