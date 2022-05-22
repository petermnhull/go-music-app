package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/pkg"
)

// GetUserByID returns individual user by ID
func GetUserByID(ctx context.Context, DB pkg.DBConnection, inputID string) (*models.User, error) {
	query := fmt.Sprintf("select * from users where id='%s'", inputID)
	row := DB.QueryRow(ctx, query)
	var (
		id           string
		accessToken  string
		refreshToken string
		scope        string
		expiresAt    time.Time
		createdAt    time.Time
		updatedAt    time.Time
	)
	err := row.Scan(
		&id,
		&accessToken,
		&refreshToken,
		&scope,
		&expiresAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, NewErrNoRecords("no matching user in database")
		}
		return nil, errors.New("failed to query database")
	}
	user := models.User{
		ID:           id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scope:        scope,
		ExpiresAt:    expiresAt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	return &user, nil
}

// UpsertUser inserts a user
func UpsertUser(ctx context.Context, DB pkg.DBConnection, user *models.User) error {
	query := fmt.Sprintf(
		`
		insert into users (
			"id",
			"access_token",
			"refresh_token",
			"scope",
			"expires_at"	
		) values ('%s', '%s', '%s', '%s', '%s')
		on conflict (id) do update set
			access_token = EXCLUDED.access_token,
			refresh_token = EXCLUDED.refresh_token,
			scope = EXCLUDED.scope,
			expires_at = EXCLUDED.expires_at;
		`,
		user.ID,
		user.AccessToken,
		user.RefreshToken,
		user.Scope,
		user.ExpiresAt.Format(time.RFC3339),
	)
	_, err := DB.Exec(ctx, query)
	return err
}
