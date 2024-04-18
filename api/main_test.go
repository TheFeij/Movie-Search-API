package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

// TestMain runs before other tests, sets gin's mode to test mode.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
