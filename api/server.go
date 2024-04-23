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
	// rapidAPI is the instance of RapidAPI for interacting with Rapid API services.
	rapidAPI rapid_api.RapidAPI
	// elasticSearch is the instance of ElasticSearch for interacting with elastic search.
	elasticSearch elastic_search.ElasticSearch
	// cache is the instance of Cacher for caching data using Redis.
	cache redis.Cacher
}

// initializeServer initializes server. Adds route handlers to the server
func initializeServer() server {
	server := server{
		router:        gin.Default(),
		rapidAPI:      rapid_api.NewRapidAPI(),
		elasticSearch: elastic_search.NewElasticSearch(),
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
	if err := server.router.Run(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
