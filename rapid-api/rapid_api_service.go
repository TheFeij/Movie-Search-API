package rapid_api

// RapidAPIService represents a service for interacting with a RapidAPI API.
type RapidAPIService interface {
	// Find searches for data based on the provided query.
	// It returns a map containing the search results and any error encountered.
	Find(query string) (map[string]interface{}, error)

	// Other methods may be added later to extend the functionality of the service.
}
