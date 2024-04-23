// Package rapid_api implements functions to interact with Rapid API's IMDb API.
package rapid_api

import "log"

// init loads configuration values
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}
