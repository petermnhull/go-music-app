package config

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// AppContext holds application level dependencies and config
type AppContext struct {
	*AppConfig
	Context      context.Context
	DBConnection DBConnectionInterface
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
	// Return context
	ctx := AppContext{
		config,
		context,
		conn,
	}
	return &ctx, nil
}

func (c *AppContext) Close() {
	_ = c.DBConnection.Close(c.Context)
}
