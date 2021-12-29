package spotify_test

import (
	"testing"

	"github.com/petermnhull/go-music-app/internal/services/spotify"
	"github.com/petermnhull/go-music-app/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	t.Run("get access ok", func(t *testing.T) {
		httpclient := testhelpers.NewMockHTTPClient()
		httpclient.AddMockResponse(
			"https://api.spotify.com/v1/me",
			200,
			`{
				"id": "123abc",
				"email": "email@email.com",
				"uri": "123abcxyz",
				"country": "GB",
				"product": "premium",
				"display_name": "Firstname Lastname",
				"product": "premium"		
			}`,
		)
		actual, err := spotify.GetMe(httpclient, "access-token")
		expected := spotify.SpotifyMe{
			ID:          "123abc",
			Email:       "email@email.com",
			URI:         "123abcxyz",
			Country:     "GB",
			Product:     "premium",
			DisplayName: "Firstname Lastname",
		}
		assert.NoError(t, err)
		assert.Equal(t, &expected, actual)
	})

	t.Run("get access validation fails", func(t *testing.T) {
		httpclient := testhelpers.NewMockHTTPClient()
		httpclient.AddMockResponse(
			"https://api.spotify.com/v1/me",
			200,
			`{"not_display_name": "123abc"}`,
		)
		_, err := spotify.GetMe(httpclient, "access-token")
		assert.Error(t, err)
		assert.Equal(t, "failed to validate profile response", err.Error())
	})
}
