// Package elastic provides functions to interact with elastic search
package elastic

import "log"

// init loads configurations
func init() {
	if err := loadConfig(); err != nil {
		log.Fatalln(err)
	}
}
