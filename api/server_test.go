package api

import (
	elasticSearchMock "Movie_Search_API/elastic-search/mock"
	rapidAPIMock "Movie_Search_API/rapid-api/mock"
	redisMock "Movie_Search_API/redis/mock"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go/scanner"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

// TestHomePage tests "/" route of the http server
func TestHomePage(t *testing.T) {
	server := newTestServer(nil, nil, nil)

	httpRequest, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	server.router.ServeHTTP(recorder, httpRequest)

	require.Equal(t, http.StatusOK, recorder.Code)

	var jsonBody map[string]interface{}

	body := recorder.Body.Bytes()
	err = json.Unmarshal(body, &jsonBody)
	require.NoError(t, err)
	require.Equal(t, "Welcome to Movie Search API!", jsonBody["message"])
}

// TestSearch tests "/search/ route of the http server
func TestSearch(t *testing.T) {

	rapidAPIHitResult := map[string]interface{}{
		"companyResults": map[string]interface{}{
			"results": []interface{}{
				map[string]interface{}{},
			},
		},
		"keywordResults": map[string]interface{}{
			"results": []interface{}{
				map[string]interface{}{},
			},
		},
		"nameResults": map[string]interface{}{
			"results": []interface{}{
				map[string]interface{}{},
			},
		},
		"titleResults": map[string]interface{}{
			"results": []interface{}{
				map[string]interface{}{},
			},
		},
	}
	rapidAPIMissResult := map[string]interface{}{
		"companyResults": map[string]interface{}{
			"results": []interface{}{},
		},
		"keywordResults": map[string]interface{}{
			"results": []interface{}{},
		},
		"nameResults": map[string]interface{}{
			"results": []interface{}{},
		},
		"titleResults": map[string]interface{}{
			"results": []interface{}{},
		},
	}
	elasticSearchMissResult := map[string]interface{}{
		"hits": map[string]interface{}{
			"hits": []interface{}{},
		},
	}
	elasticSearchHitResult := map[string]interface{}{
		"hits": map[string]interface{}{
			"hits": []interface{}{
				map[string]interface{}{},
			},
		},
	}
	cacheHitResult := rapidAPIHitResult // or elasticSearchHitResult

	testCases := []struct {
		name       string
		query      string
		buildStubs func(
			cache *redisMock.MockCacheService,
			rapidAPI *rapidAPIMock.MockRapidAPIService,
			elasticSearch *elasticSearchMock.MockElasticSearchService,
			query string)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "NotFound",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(rapidAPIMissResult, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "GetFromRapidAPI",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().Find(query).Times(1).Return(rapidAPIHitResult, nil)

				cache.EXPECT().
					SetData("search:"+query, rapidAPIHitResult, 24*time.Hour).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				body, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				require.NoError(t, err)

				require.Equal(t, rapidAPIHitResult, result)
			},
		},
		{
			name:  "GetFromElasticSearch",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchHitResult, nil)

				cache.EXPECT().
					SetData("search:"+query, elasticSearchHitResult, 24*time.Hour).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				body, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				require.NoError(t, err)

				require.Equal(t, elasticSearchHitResult, result)
			},
		},
		{
			name:  "GetFromCache",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(cacheHitResult, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				body, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				require.NoError(t, err)

				require.Equal(t, cacheHitResult, result)
			},
		},
		{
			name:  "ElasticSearchUrlError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(map[string]interface{}{}, url.Error{Err: fmt.Errorf("error for testing purposes")}.Err)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "ElasticSearchJsonUnmarshalError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(map[string]interface{}{}, scanner.Error{Msg: "error for testing purposes"})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "RapidAPIInvalidMethodError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(map[string]interface{}{}, fmt.Errorf("net/http: invalid method"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "RapidAPIUrlError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(map[string]interface{}{}, url.Error{Err: fmt.Errorf("error for testing purposes")}.Err)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "RapidAPIJsonUnmarshalError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(map[string]interface{}{}, scanner.Error{Msg: "error for testing purposes"})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "RapidAPIStatusNotOK",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(map[string]interface{}{}, fmt.Errorf("status code: not 200"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "RapidAPIEOFError",
			query: "Tom Cruise",
			buildStubs: func(
				cache *redisMock.MockCacheService,
				rapidAPI *rapidAPIMock.MockRapidAPIService,
				elasticSearch *elasticSearchMock.MockElasticSearchService,
				query string,
			) {
				cache.EXPECT().
					GetData("search:"+query).
					Times(1).
					Return(map[string]interface{}{}, redis.Nil)

				elasticSearch.EXPECT().
					SearchQuery(query).
					Times(1).
					Return(elasticSearchMissResult, nil)

				rapidAPI.EXPECT().
					Find(query).
					Times(1).
					Return(map[string]interface{}{}, io.EOF)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			cache := redisMock.NewMockCacheService(controller)
			rapidAPI := rapidAPIMock.NewMockRapidAPIService(controller)
			elasticSearch := elasticSearchMock.NewMockElasticSearchService(controller)

			testCase.buildStubs(cache, rapidAPI, elasticSearch, testCase.query)

			httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search?query=%s", url.QueryEscape(testCase.query)), nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			server := newTestServer(rapidAPI, elasticSearch, cache)
			server.router.ServeHTTP(recorder, httpReq)

			testCase.checkResponse(t, recorder)
		})
	}
}
