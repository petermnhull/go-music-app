package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig stores application configuration
type AppConfig struct {
	Port         int           `envconfig:"APP_PORT" default:"8080"`
	Address      string        `envconfig:"APP_ADDRESS" default:"127.0.0.1"`
	ReadTimeout  time.Duration `envconfig:"APP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"APP_WRITE_TIMEOUT" default:"5s"`
	DatabaseURL  string        `envconfig:"DATABASE_URL"`
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
