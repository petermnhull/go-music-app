package endpoints

import (
	"net/http"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/rs/zerolog/log"
)

// HealthCheckHandler exposes health check endpoint
func HealthCheckHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	log.Info().Msg("Health check endpoint hit!")
	data := map[string]string{
		"user_agent": r.Header.Get("User-Agent"),
	}
	response := NewAPIResponseSuccess(http.StatusOK, data)
	return response
}
