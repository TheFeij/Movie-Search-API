// Package rapid_api implements functions to interact with Rapid API's IMDb API.
package rapid_api

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
