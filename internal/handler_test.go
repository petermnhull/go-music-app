package internal_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServeHTTP(t *testing.T) {
	t.Run("handler wrapper handles health check prints ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
			pgxmock.MonitorPingsOption(true),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectPing()

		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		request.Header.Add("User-Agent", "Go Test Suite")
		response := httptest.NewRecorder()
		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}
		ah := internal.AppHandler{&ctx, endpoints.HealthCheckHandler}

		ah.ServeHTTP(response, request)
		actual := response.Body.String()
		expected := `{"code": 200, "status": "success", "data": {"user_agent": "Go Test Suite"}}`

		require.JSONEq(t, expected, actual)
		assert.Equal(t, response.Result().StatusCode, 200)
	})
}
