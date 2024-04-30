// Package redis provides functions to set json data to redis and get json data from it
package redis

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
