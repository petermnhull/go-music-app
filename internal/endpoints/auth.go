package endpoints

import (
	"net/http"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/services/spotify"
)

// AuthHandler continues authorization flow by accepting an access code
func AuthHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	var err error

	// Retrieve code
	code := r.URL.Query().Get("code")
	if code == "" {
		return NewAPIResponseFailed(http.StatusBadRequest, "missing access code")
	}

	// Obtain access and refresh tokens.
	access, err := spotify.GetAccess(
		ctx.HTTPClient,
		code,
		ctx.AppConfig.SpotifyRedirectURI,
		ctx.AppConfig.SpotifyClientID,
		ctx.AppConfig.SpotifyClientSecret,
	)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, err.Error())
	}

	// Find out who it is
	user, err := spotify.GetMe(ctx.HTTPClient, access.AccessToken)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, err.Error())
	}

	// TODO: Store data in DB

	return NewAPIResponseSuccess(http.StatusOK, user)
}
