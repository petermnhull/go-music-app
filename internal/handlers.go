package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type apiResponse struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// HealthCheckHandler exposes health check endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Health check endpoint hit!")
	w.WriteHeader(http.StatusOK)
	data := map[string]string{}
	response := apiResponse{200, "success", data}
	responseBytes, _ := json.Marshal(response)
	fmt.Fprint(w, string(responseBytes))
}
