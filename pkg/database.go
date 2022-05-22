package pkg

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// DBConnection provides generic interface for interacting with database
type DBConnection interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}
