package config

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type DBConnectionInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}
