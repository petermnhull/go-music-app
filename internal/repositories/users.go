package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/models"
)

// CountUsers returns a count of all the users
func CountUsers(ctx context.Context, DB config.DBConnectionInterface) (int64, error) {
	row := DB.QueryRow(ctx, "select count(*) from users")
	var count int64
	err := row.Scan(&count)
	return count, err
}

// GetAllUsers returns all users
func GetAllUsers(ctx context.Context, DB config.DBConnectionInterface) ([]models.User, error) {
	rows, err := DB.Query(ctx, "select * from users")
	if err != nil {
		return nil, errors.New("failed to query user data")
	}

	users := []models.User{}
	for rows.Next() {
		var (
			id              int64
			spotifyUsername string
			lastfmUsername  string
			createdAt       time.Time
			updatedAt       time.Time
		)
		err := rows.Scan(&id, &spotifyUsername, &lastfmUsername, &createdAt, &updatedAt)
		if err != nil {
			return nil, errors.New("failed to scan row")
		}
		user := models.User{
			ID:              id,
			SpotifyUsername: spotifyUsername,
			LastfmUsername:  lastfmUsername,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}

// GetUserByID returns individual user by ID
func GetUserByID(ctx context.Context, DB config.DBConnectionInterface, userID int64) (*models.User, error) {
	query := fmt.Sprintf("select * from users where id=%d", userID)
	row := DB.QueryRow(ctx, query)
	var (
		id              int64
		spotifyUsername string
		lastfmUsername  string
		createdAt       time.Time
		updatedAt       time.Time
	)
	err := row.Scan(&id, &spotifyUsername, &lastfmUsername, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &ErrNoRecords{"no matching user in database"}
		}
		return nil, errors.New("failed to query database")
	}
	user := models.User{
		ID:              id,
		SpotifyUsername: spotifyUsername,
		LastfmUsername:  lastfmUsername,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	return &user, nil
}
