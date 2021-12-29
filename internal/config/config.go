package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig stores application configuration
type AppConfig struct {
	ServerConfig
	DatabaseConfig
	HTTPClientConfig
	SpotifyConfig
	ClientConfig
}

// ServerConfig for base settings for the app
type ServerConfig struct {
	Name         string        `envconfig:"APP_NAME" default:"go-music-app"`
	Port         int           `envconfig:"APP_PORT" default:"8080"`
	Address      string        `envconfig:"APP_ADDRESS" default:"127.0.0.1"`
	Protocol     string        `envconfig:"APP_PROTOCOL" default:"http"`
	ReadTimeout  time.Duration `envconfig:"APP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"APP_WRITE_TIMEOUT" default:"5s"`
}

// DatabaseConfig for connecting to the Postgres database
type DatabaseConfig struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

// HTTPClientConfig for HTTP requests
type HTTPClientConfig struct {
	HTTPClientTimeout           time.Duration `envconfig:"HTTP_CLIENT_TIMEOUT" default:"5s"`
	HTTPClientBackoffInternal   time.Duration `envconfig:"HTTP_CLIENT_BACKOFF_INTERVAL" default:"2ms"`
	HTTPClientMaxJitterInternal time.Duration `envconfig:"HTTP_CLIENT_MAX_JITTER_INTERNAL" default:"5ms"`
}

// SpotifyConfig for calling Spotify API
type SpotifyConfig struct {
	SpotifyClientID     string `envconfig:"SPOTIFY_CLIENT_ID" required:"true"`
	SpotifyClientSecret string `envconfig:"SPOTIFY_CLIENT_SECRET" required:"true"`
	SpotifyRedirectURI  string `envconfig:"SPOTIFY_REDIRECT_URI" required:"true"`
}

// ClientConfig for front end specific config
type ClientConfig struct {
	AllowOriginURI string `envconfig:"ALLOW_ORIGIN_URI" required:"true"`
}

// newAppConfig reads configuration from environment variables
func newAppConfig() (*AppConfig, error) {
	var a AppConfig
	err := envconfig.Process("", &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
