package spotify_test

import (
	"testing"

	"github.com/petermnhull/go-music-app/internal/services/spotify"
	"github.com/petermnhull/go-music-app/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGetAccess(t *testing.T) {
	t.Run("get access ok", func(t *testing.T) {
		httpclient := testhelpers.NewMockHTTPClient()
		httpclient.AddMockResponse(
			"https://accounts.spotify.com/api/token",
			200,
			`{"access_token": "123abc", "refresh_token": "789xyz", "scope": "read-private", "token_type": "token"}`,
		)
		actual, err := spotify.GetAccess(httpclient, "code", "redirect", "clientID", "clientSecret")
		expected := spotify.SpotifyAccess{
			AccessToken:  "123abc",
			RefreshToken: "789xyz",
			Scope:        "read-private",
			TokenType:    "token",
		}
		assert.NoError(t, err)
		assert.Equal(t, &expected, actual)
	})

	t.Run("get access fails", func(t *testing.T) {
		httpclient := testhelpers.NewMockHTTPClient()
		httpclient.AddMockResponse(
			"https://accounts.spotify.com/api/token",
			400,
			`{"error": "missing access code"}`,
		)
		actual, err := spotify.GetAccess(httpclient, "", "redirect", "clientID", "clientSecret")
		assert.Error(t, err)
		assert.Equal(t, "external request to spotify failed (400): failed to retrieve tokens", err.Error())
		assert.Nil(t, actual)
	})
}
