// Code generated by MockGen. DO NOT EDIT.
// Source: Movie_Search_API/elastic-search (interfaces: ElasticSearchService)
//
// Generated by this command:
//
//	mockgen -package mock -destination elastic-search/mock/elastic_search_services.go Movie_Search_API/elastic-search ElasticSearchService
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockElasticSearchService is a mock of ElasticSearchService interface.
type MockElasticSearchService struct {
	ctrl     *gomock.Controller
	recorder *MockElasticSearchServiceMockRecorder
}

// MockElasticSearchServiceMockRecorder is the mock recorder for MockElasticSearchService.
type MockElasticSearchServiceMockRecorder struct {
	mock *MockElasticSearchService
}

// NewMockElasticSearchService creates a new mock instance.
func NewMockElasticSearchService(ctrl *gomock.Controller) *MockElasticSearchService {
	mock := &MockElasticSearchService{ctrl: ctrl}
	mock.recorder = &MockElasticSearchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElasticSearchService) EXPECT() *MockElasticSearchServiceMockRecorder {
	return m.recorder
}

// SearchQuery mocks base method.
func (m *MockElasticSearchService) SearchQuery(arg0 string) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchQuery", arg0)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchQuery indicates an expected call of SearchQuery.
func (mr *MockElasticSearchServiceMockRecorder) SearchQuery(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuery", reflect.TypeOf((*MockElasticSearchService)(nil).SearchQuery), arg0)
}
