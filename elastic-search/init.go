// Package elastic_search provides functions to interact with elastic-search search
package elastic_search

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
