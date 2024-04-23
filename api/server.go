// Package api implements an HTTP API server.
//
// It provides the following APIs:
//
// - "/" -> home page
package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// server holds a router value of type *gin.Engine
var server struct {
	router *gin.Engine
}

// init loads configurations
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}

// initializeServer initializes server. Adds route handlers to the server
func initializeServer() {
	server.router = gin.Default()

	server.router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Welcome to Movie Search API!"})
	})

	server.router.GET("/search", search)
}

// StartServer starts the server on the address specified in the configuration file
func StartServer() {
	initializeServer()
	if err := server.router.Run(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
