// Package redis provides functions to set json data to redis and get json data from it
package redis

import "log"

// init loads configuration values
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}
