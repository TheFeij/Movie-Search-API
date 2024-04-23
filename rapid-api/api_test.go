package rapid_api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestFind tests Find
func TestFind(t *testing.T) {
	rapidAPIService := NewRapidAPIService()

	result, err := rapidAPIService.Find("brad")
	require.NoError(t, err)

	require.NotEmpty(t, result)
}
