package config

import (
	"context"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4"
	"github.com/petermnhull/go-music-app/pkg"
)

// AppContext holds application level dependencies and config
type AppContext struct {
	*AppConfig
	Context      context.Context
	HTTPClient   pkg.HTTPClient
	DBConnection pkg.DBConnection
}

// NewContext creates a new context from environment variables
func NewContext() (*AppContext, error) {
	// Load config
	config, err := newAppConfig()
	if err != nil {
		return nil, err
	}

	// Base context
	context := context.Background()

	// Database connection
	conn, err := pgx.Connect(context, config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// HTTP client
	backoff := heimdall.NewConstantBackoff(
		config.HTTPClientBackoffInternal,
		config.HTTPClientMaxJitterInternal,
	)
	retrier := heimdall.NewRetrier(backoff)
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(config.HTTPClientTimeout),
		httpclient.WithRetrier(retrier),
	)

	// Return context
	ctx := AppContext{
		config,
		context,
		client,
		conn,
	}
	return &ctx, nil
}

func (c *AppContext) Close() {
	_ = c.DBConnection.Close(c.Context)
}
