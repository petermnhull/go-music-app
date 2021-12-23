package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/petermnhull/go-music-app/internal"
	"github.com/rs/zerolog/log"
)

func main() {
	port := os.Getenv("APP_PORT")

	log.Info().Msg(fmt.Sprintf("Starting server on port %s", port))

	server := &http.Server{
		Handler:      internal.NewRouter(),
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal().Err(server.ListenAndServe())
}
