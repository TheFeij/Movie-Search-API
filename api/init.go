// Package api implements an HTTP API server.
//
// It provides the following APIs:
//
// - "/" -> home page
// - "/search" -> search user's query
package api

import (
	"Movie_Search_API/config"
	"log"
)

var configurations config.Config

// init loads configurations
func init() {
	var err error
	configurations, err = config.LoadConfig("./config", "config")
	if err != nil {
		log.Fatalln(err)
	}
}
