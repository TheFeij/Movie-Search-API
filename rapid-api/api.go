package rapid_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	if err := loadConfig(); err != nil {
		return
	}
}

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
