test:
	go test -v -cover ./...

godoc:
	godoc -http=:6060

mock:
	$(shell cd elastic-search && mockgen -package mock -destination mock/elastic_search_services.go Movie_Search_API/elastic-search ElasticSearchService)
	$(shell cd rapid-api && mockgen -package mock -destination mock/rapi_api_services.go Movie_Search_API/rapid-api RapidAPIService)
	$(shell cd redis && mockgen -package mock -destination mock/cache_service.go Movie_Search_API/redis CacheService)