package api

import (
	"Movie_Search_API/elastic-search"
	"Movie_Search_API/rapid-api"
	"Movie_Search_API/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func search(context *gin.Context) {
	query := context.Query("query")

	cacheKey := "search:" + query
	jsonData, err := redis.GetData(cacheKey)
	if err == nil {
		context.JSON(http.StatusOK, jsonData)
		return
	}

	jsonData, err = elastic_search.SearchQuery(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	hits := len(jsonData["hits"].(map[string]interface{})["hits"].([]interface{}))
	if hits > 0 {
		cacheResult(query, jsonData)
		context.JSON(http.StatusOK, jsonData)
		return
	}

	jsonData, err = rapid_api.SearchQuery(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	if len(jsonData["companyResults"].(map[string]interface{})["results"].([]interface{})) == 0 &&
		len(jsonData["keywordResults"].(map[string]interface{})["results"].([]interface{})) == 0 &&
		len(jsonData["nameResults"].(map[string]interface{})["results"].([]interface{})) == 0 &&
		len(jsonData["titleResults"].(map[string]interface{})["results"].([]interface{})) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "No results found"})
		return
	}

	cacheResult(query, jsonData)
	context.JSON(http.StatusOK, jsonData)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func cacheResult(query string, jsonData map[string]interface{}) {
	cacheKey := "search:" + query
	err := redis.SetData(cacheKey, jsonData, 24*time.Hour)
	if err != nil {
		log.Println("Cache error:", err)
	}
}
