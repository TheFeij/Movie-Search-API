// Package redis provides functions to set json data to redis and get json data from it
package redis

// init loads configuration values
func init() {
	if err := loadConfig(); err != nil {
		return
	}
}
