package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/rs/zerolog/log"
)

// HealthCheckHandler exposes health check endpoint
func HealthCheckHandler(ctx *config.AppContext, w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Health check endpoint hit!")
	w.WriteHeader(http.StatusOK)
	data := map[string]string{}
	response := apiResponse{200, "success", data}
	responseBytes, _ := json.Marshal(response)
	fmt.Fprint(w, string(responseBytes))
}
