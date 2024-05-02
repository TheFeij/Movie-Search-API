test:
	go test -v -cover ./...

godoc:
	godoc -http=:6060

mock:
	mockgen -package mock -destination elastic-search/mock/elastic_search_services.go Movie_Search_API/elastic-search ElasticSearchService
	mockgen -package mock -destination rapid-api/mock/rapi_api_services.go Movie_Search_API/rapid-api RapidAPIService
	mockgen -package mock -destination redis/mock/cache_service.go Movie_Search_API/redis CacheService

deploy-kubernetes:
	kubectl apply -f kubernetes/

undeploy-kubernetes:
	kubectl delete -f kubernetes/