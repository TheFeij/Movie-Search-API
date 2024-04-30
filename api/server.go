package api

import (
	elastic_search "Movie_Search_API/elastic-search"
	rapid_api "Movie_Search_API/rapid-api"
	"Movie_Search_API/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type server struct {
	// router is the Gin router instance used for handling HTTP requests.
	router *gin.Engine
	// rapidAPI is the service used for interacting with RapidAPI.
	rapidAPI rapid_api.RapidAPIService
	// elasticSearch is the service used for interacting with elastic search.
	elasticSearch elastic_search.ElasticSearchService
	// cache is the instance of CacheService for caching data using Redis.
	cache redis.CacheService
}

// initializeServer initializes server. Adds route handlers to the server
func initializeServer() server {
	server := server{
		router:        gin.Default(),
		rapidAPI:      rapid_api.NewRapidAPIService(),
		elasticSearch: elastic_search.NewElasticSearchService(),
		cache:         redis.GetCacher(),
	}

	server.router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Welcome to Movie Search API!"})
	})

	server.router.GET("/search", server.search)

	return server
}

// StartServer starts the server on the address specified in the configuration file
func StartServer() {
	server := initializeServer()
	if err := server.router.Run(configurations.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
