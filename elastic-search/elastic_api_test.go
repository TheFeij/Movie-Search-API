package elastic_search

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchQuery(t *testing.T) {
	elasticSearchService := NewElasticSearchService()

	t.Run("OK", func(t *testing.T) {
		query := "Tom Cruise"

		result, err := elasticSearchService.SearchQuery(query)
		require.NoError(t, err)
		require.NotEmpty(t, result)

		fmt.Println(result)

		hitsMap, ok := result["hits"]
		require.True(t, ok, "hits key not found in result")
		require.NotNil(t, hitsMap, "hits value is nil")

		hitsArray, ok := result["hits"].(map[string]interface{})["hits"]
		require.True(t, ok, "hits.hits key not found in result")
		require.NotNil(t, hitsArray, "hits.hits value is nil")

		totalMap, ok := result["hits"].(map[string]interface{})["total"]
		require.True(t, ok, "hits.total key not found in result")
		require.NotNil(t, totalMap, "hits.total value is nil")

		timedOut, ok := result["timed_out"]
		require.True(t, ok, "hits.timed_out key not found in result")
		require.False(t, timedOut.(bool))

		took, ok := result["took"]
		require.True(t, ok, "hits.took key not found in result")
		require.NotNil(t, took, "hits.took key not found in result")
	})
}
