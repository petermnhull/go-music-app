package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	t.Run("health check ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		request.Header.Add("User-Agent", "Go Test Suite")
		ctx := config.AppContext{}

		actual := endpoints.HealthCheckHandler(&ctx, request)
		expected := endpoints.APIResponse{200, "success", map[string]string{"user_agent": "Go Test Suite"}}
		assert.Equal(t, &expected, actual)
	})
}
