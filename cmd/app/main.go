package main

import (
	"fmt"
	"net/http"

	"github.com/petermnhull/go-music-app/internal"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/rs/zerolog/log"
)

func buildServer(ctx *config.AppContext) *http.Server {
	router := internal.NewRouter(ctx)
	address := fmt.Sprintf("%s:%d", ctx.AppConfig.Address, ctx.AppConfig.Port)
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: ctx.AppConfig.ReadTimeout,
		ReadTimeout:  ctx.AppConfig.WriteTimeout,
	}
	return server
}

func main() {
	ctx, err := config.NewContext()
	if err != nil {
		log.Fatal().Msg("failed to load config: " + err.Error())
	}
	defer ctx.Close()

	server := buildServer(ctx)

	log.Info().Msg(fmt.Sprintf("Starting server on port %v", ctx.AppConfig.Port))
	log.Fatal().Err(server.ListenAndServe())
}
