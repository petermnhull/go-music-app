package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petermnhull/go-music-app/internal"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckEndpoint(t *testing.T) {
	t.Run("health check prints ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		internal.HealthCheckHandler(response, request)

		actual := response.Body.String()
		expected := `{"code": 200,"status": "success","data": {}}`

		require.JSONEq(t, actual, expected, "should be equal")
	})
}
