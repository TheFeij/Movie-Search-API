package rapid_api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSearchQuery tests SearchQuery
func TestSearchQuery(t *testing.T) {
	searcher := NewSearcher()

	result, err := searcher.SearchQuery("brad")
	require.NoError(t, err)

	require.NotEmpty(t, result)
}
