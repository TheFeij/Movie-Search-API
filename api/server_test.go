package api

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomePage tests "/" route of the http server
func TestHomePage(t *testing.T) {
	server := initializeServer()

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
