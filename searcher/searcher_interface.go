// Package searcher contains Searcher Interface
package searcher

// Searcher  an interface for performing search operations.
type Searcher interface {
	// SearchQuery searches for the provided query and returns the result as a map[string]interface{}.
	// It returns an error if the search operation fails.
	SearchQuery(query string) (map[string]interface{}, error)
}
