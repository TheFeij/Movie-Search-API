// Package api implements an HTTP API server.
//
// It provides the following APIs:
//
// - "/" -> home page
// - "/search" -> search user's query
package api

import "log"

// init loads configurations
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}
