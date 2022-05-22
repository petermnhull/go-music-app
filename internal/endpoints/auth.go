package endpoints

import (
	"net/http"
	"time"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/internal/repositories"
	"github.com/petermnhull/go-music-app/pkg/spotify"
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
	profile, err := spotify.GetMe(ctx.HTTPClient, access.AccessToken)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, err.Error())
	}

	// Upsert latest tokens into database
	expiresAt := time.Now().Add(time.Second * time.Duration(access.ExpiresIn)).UTC()
	user := &models.User{
		ID:           profile.ID,
		AccessToken:  access.AccessToken,
		RefreshToken: access.RefreshToken,
		Scope:        access.Scope,
		ExpiresAt:    expiresAt,
	}
	err = repositories.UpsertUser(ctx.Context, ctx.DBConnection, user)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, err.Error())
	}

	return NewAPIResponseSuccess(http.StatusOK, user)
}
