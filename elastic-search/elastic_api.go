package elastic_search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// SearchQuery receives a query, calls elastic-search search api to search for documents related tot the query
func SearchQuery(query string) (map[string]interface{}, error) {
	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s/movies/_search?q=%s", config.ElasticSearchAddress, encodedQuery)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	return jsonData, err
}
