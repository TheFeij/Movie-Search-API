package api

import (
	elastic_search "Movie_Search_API/elastic-search"
	rapid_api "Movie_Search_API/rapid-api"
	"Movie_Search_API/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"testing"
)

// initializeServer initializes server. Adds route handlers to the server
func newTestServer(
	rapidAPI rapid_api.RapidAPIService,
	elasticSearch elastic_search.ElasticSearchService,
	cache redis.CacheService,
) server {

	server := server{
		router:        gin.Default(),
		rapidAPI:      rapidAPI,
		elasticSearch: elasticSearch,
		cache:         cache,
	}

	server.router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Welcome to Movie Search API!"})
	})

	server.router.GET("/search", server.search)

	return server
}

// TestMain runs before other tests, sets gin's mode to test mode.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
