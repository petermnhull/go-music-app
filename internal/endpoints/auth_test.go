package endpoints_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/petermnhull/go-music-app/internal/services/spotify"
	"github.com/petermnhull/go-music-app/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler(t *testing.T) {
	t.Run("auth 400", func(t *testing.T) {
		ctx := config.AppContext{}
		request, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		actual := endpoints.AuthHandler(&ctx, request)
		expected := endpoints.APIResponse{400, "failed", map[string]string{"error": "missing access code"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("auth 404", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		httpclient := testhelpers.NewMockHTTPClient()

		appConfig := config.AppConfig{
			SpotifyConfig: config.SpotifyConfig{
				SpotifyRedirectURI:  "redirect",
				SpotifyClientID:     "client-id",
				SpotifyClientSecret: "client-secret",
			},
		}

		ctx := config.AppContext{
			AppConfig:    &appConfig,
			Context:      context.Background(),
			DBConnection: mockDB,
			HTTPClient:   httpclient,
		}

		request, _ := http.NewRequest(http.MethodGet, "/auth?code=123abc", nil)
		actual := endpoints.AuthHandler(&ctx, request)
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "external request to spotify failed (404): failed to retrieve tokens"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("auth ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		httpclient := testhelpers.NewMockHTTPClient()
		httpclient.AddMockResponse(
			"https://accounts.spotify.com/api/token",
			200,
			`{"access_token": "123abc", "refresh_token": "789xyz", "scope": "read-private", "token_type": "token"}`,
		)
		httpclient.AddMockResponse(
			"https://api.spotify.com/v1/me",
			200,
			`{"display_name": "Firstname Lastname"}`,
		)

		appConfig := config.AppConfig{
			SpotifyConfig: config.SpotifyConfig{
				SpotifyRedirectURI:  "redirect",
				SpotifyClientID:     "client-id",
				SpotifyClientSecret: "client-secret",
			},
		}

		ctx := config.AppContext{
			AppConfig:    &appConfig,
			Context:      context.Background(),
			DBConnection: mockDB,
			HTTPClient:   httpclient,
		}

		request, _ := http.NewRequest(http.MethodGet, "/auth?code=999yyy", nil)
		actual := endpoints.AuthHandler(&ctx, request)
		expected := endpoints.APIResponse{200, "success", &spotify.SpotifyMe{DisplayName: "Firstname Lastname"}}
		assert.Equal(t, &expected, actual)
	})
}
