package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCountUsers(t *testing.T) {
	t.Run("count users ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		columns := []string{"count"}
		mockDB.ExpectQuery(
			"select count(*) from users",
		).WillReturnRows(
			mockDB.NewRows(columns).AddRow(int64(1)),
		)
		result, err := repositories.CountUsers(context.Background(), mockDB)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("get all users ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		columns := []string{"id", "spotify_username", "lastfm_username", "created_at", "updated_at"}
		time := time.Now().UTC()
		mockDB.ExpectQuery(
			"select * from users",
		).WillReturnRows(
			mockDB.NewRows(columns).AddRow(int64(1), "username1", "username2", time, time),
		)

		result, err := repositories.GetAllUsers(context.Background(), mockDB)
		assert.NoError(t, err)
		user := models.User{
			ID:              1,
			SpotifyUsername: "username1",
			LastfmUsername:  "username2",
			CreatedAt:       time,
			UpdatedAt:       time,
		}
		expected := []models.User{user}
		assert.Equal(t, expected, result)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("get user ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		columns := []string{"id", "spotify_username", "lastfm_username", "created_at", "updated_at"}
		time := time.Now().UTC()
		mockDB.ExpectQuery(
			"select * from users where id=1",
		).WillReturnRows(
			mockDB.NewRows(columns).AddRow(int64(1), "username1", "username2", time, time),
		)

		result, err := repositories.GetUserByID(context.Background(), mockDB, 1)
		assert.NoError(t, err)
		expected := models.User{
			ID:              1,
			SpotifyUsername: "username1",
			LastfmUsername:  "username2",
			CreatedAt:       time,
			UpdatedAt:       time,
		}
		assert.Equal(t, &expected, result)
	})

	t.Run("get user not found", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		mockDB.ExpectQuery("select * from users where id=1").WillReturnError(pgx.ErrNoRows)

		_, err = repositories.GetUserByID(context.Background(), mockDB, 1)
		assert.Error(t, err)
		assert.EqualError(t, err, "no matching user in database")
	})
}
