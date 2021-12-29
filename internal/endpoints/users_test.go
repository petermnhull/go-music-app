package endpoints_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDHandler(t *testing.T) {
	t.Run("get user fails with 404", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectQuery("select * from users where id='123abc'").WillReturnError(pgx.ErrNoRows)

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/123abc", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "123abc"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expected := endpoints.APIResponse{404, "failed", map[string]string{"error": "no matching user in database"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("get user fails with 500", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/123abc", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "123abc"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "failed to retrieve users: failed to query database"}}
		assert.Equal(t, &expected, actual)
	})

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
			"scopes",
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

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/123abc", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "123abc"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expectedUser := models.User{
			ID:           "123abc",
			AccessToken:  "access-token",
			RefreshToken: "refresh-token",
			Scope:        "read-private",
			ExpiresAt:    time,
			CreatedAt:    time,
			UpdatedAt:    time,
		}
		expected := endpoints.APIResponse{200, "success", map[string]models.User{"user": expectedUser}}
		assert.Equal(t, &expected, actual)
	})
}
