package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	t.Run("get user ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		columns := []string{
			"id",
			"access_token",
			"refresh_token",
			"scope",
			"expires_at",
			"created_at",
			"updated_at",
		}
		time := time.Now().UTC()
		mockDB.ExpectQuery(
			"select * from users where id='123abc'",
		).WillReturnRows(
			mockDB.NewRows(columns).AddRow(
				"123abc",
				"access-token",
				"refresh-token",
				"read-private",
				time,
				time,
				time,
			),
		)

		result, err := repositories.GetUserByID(context.Background(), mockDB, "123abc")
		assert.NoError(t, err)
		expected := models.User{
			ID:           "123abc",
			AccessToken:  "access-token",
			RefreshToken: "refresh-token",
			Scope:        "read-private",
			ExpiresAt:    time,
			CreatedAt:    time,
			UpdatedAt:    time,
		}
		assert.Equal(t, &expected, result)
	})

	t.Run("get user not found", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		mockDB.ExpectQuery("select * from users where id='123abc'").WillReturnError(pgx.ErrNoRows)

		_, err = repositories.GetUserByID(context.Background(), mockDB, "123abc")
		assert.Error(t, err)
		assert.EqualError(t, err, "no matching user in database")
	})
}

func TestUpsertUser(t *testing.T) {
	t.Run("upsert user ok", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		query := `insert into users (
			"id",
			"access_token",
			"refresh_token",
			"scope",
			"expires_at"
		) values
		('123abc', 'access-token', 'refresh-token', 'read-private', '2020-01-01T00:00:00Z')
		on conflict (id) do update set
		access_token = EXCLUDED.access_token,
		refresh_token = EXCLUDED.refresh_token,
		scope = EXCLUDED.scope,
		expires_at = EXCLUDED.expires_at;
		`
		mockDB.ExpectExec(query).WillReturnResult(pgconn.CommandTag{})
		user := models.User{
			ID:           "123abc",
			AccessToken:  "access-token",
			RefreshToken: "refresh-token",
			Scope:        "read-private",
			ExpiresAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		err = repositories.UpsertUser(context.Background(), mockDB, &user)
		assert.NoError(t, err)
	})
}
