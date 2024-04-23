package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (s server) search(context *gin.Context) {
	query := context.Query("query")

	cacheKey := "search:" + query
	jsonData, err := s.cache.GetData(cacheKey)
	if err == nil {
		context.JSON(http.StatusOK, jsonData)
		return
	}

	jsonData, err = s.elasticSearch.SearchQuery(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	hits := len(jsonData["hits"].(map[string]interface{})["hits"].([]interface{}))
	if hits > 0 {
		s.cacheResult(query, jsonData)
		context.JSON(http.StatusOK, jsonData)
		return
	}

	jsonData, err = s.rapidAPI.Find(query)
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

	s.cacheResult(query, jsonData)
	context.JSON(http.StatusOK, jsonData)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s server) cacheResult(query string, jsonData map[string]interface{}) {
	cacheKey := "search:" + query
	err := s.cache.SetData(cacheKey, jsonData, 24*time.Hour)
	if err != nil {
		log.Println("Cache error:", err)
	}
}
