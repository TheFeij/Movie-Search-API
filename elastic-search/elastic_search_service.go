package elastic_search

// ElasticSearchService represents a service for interacting with elastic search.
type ElasticSearchService interface {
	// SearchQuery searches for the provided query and returns the result as a map[string]interface{}.
	// It returns an error if the search operation fails.
	SearchQuery(query string) (map[string]interface{}, error)
}
