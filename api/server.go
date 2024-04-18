package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var server struct {
	router *gin.Engine
}

func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}

func initializeServer() {
	server.router = gin.Default()

	server.router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Welcome to Movie Search API!"})
	})

	//TODO: /search GET endpoint
}

func StartServer() {
	initializeServer()
	if err := server.router.Run(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
