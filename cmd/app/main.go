package main

import (
	"fmt"

	"github.com/petermnhull/go-music-app/internal"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, err := config.NewContext()
	if err != nil {
		log.Fatal().Msg("failed to load config: " + err.Error())
	}
	defer ctx.Close()
	server := internal.NewServer(ctx)
	log.Info().Msg(fmt.Sprintf("Starting server on port %v", ctx.AppConfig.Port))
	log.Fatal().Err(server.ListenAndServe())
}
