// Package rapid_api implements functions to interact with Rapid API's IMDb API.
package rapid_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// init loads configuration values
func init() {
	if err := loadConfig(); err != nil {
		return
	}
}

// SearchQuery performs a search query using Rapid API's IMDb API.
func SearchQuery(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s=%s", config.RapidAPISearchEndpoint, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return make(map[string]interface{}), err
	}

	req.Header.Add("X-RapidAPI-Key", config.XRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", config.XRapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return make(map[string]interface{}), err
	}
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("status code: %v\n", res.StatusCode)
		return make(map[string]interface{}), err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return make(map[string]interface{}), err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)

	return jsonData, err
}
