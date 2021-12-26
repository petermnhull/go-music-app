package internal_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	t.Run("server set up ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn()
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectPing()

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		router := internal.NewRouter(&ctx)
		server := httptest.NewServer(router)
		defer server.Close()

		url := server.URL + "/health"
		r, err := http.Get(url)
		assert.NoError(t, err)
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		expected := `{"code": 200, "status": "success", "data": {"user_agent": "Go-http-client/1.1"}}`
		require.JSONEq(t, expected, string(body))
	})
}
