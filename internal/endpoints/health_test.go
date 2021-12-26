package endpoints_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	t.Run("health check ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
			pgxmock.MonitorPingsOption(true),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectPing()

		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		request.Header.Add("User-Agent", "Go Test Suite")
		context := context.Background()
		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context,
			DBConnection: mockDB,
		}

		actual := endpoints.HealthCheckHandler(&ctx, request)
		expected := endpoints.APIResponse{200, "success", map[string]string{"user_agent": "Go Test Suite"}}
		assert.Equal(t, &expected, actual)
	})
}
