// Package elastic_search provides functions to interact with elastic-search search
package elastic_search

import "log"

// init loads configurations
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}
