// Package rapid_api implements functions to interact with Rapid API's IMDb API.
package rapid_api

// init loads configuration values
func init() {
	if err := loadConfig(); err != nil {
		return
	}
}
