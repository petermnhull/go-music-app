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

func TestUsersCountHandler(t *testing.T) {
	t.Run("users count fails", func(t *testing.T) {
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

		request, _ := http.NewRequest(http.MethodGet, "/users/count", nil)
		actual := endpoints.UsersCountHandler(&ctx, request)
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "failed to get count of users"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("users count ok", func(t *testing.T) {
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

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		request, _ := http.NewRequest(http.MethodGet, "/users/count", nil)
		actual := endpoints.UsersCountHandler(&ctx, request)
		expected := endpoints.APIResponse{200, "success", map[string]string{"count": "1"}}
		assert.Equal(t, &expected, actual)
	})
}
func TestGetUsersHandler(t *testing.T) {
	t.Run("get users fails", func(t *testing.T) {
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

		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		actual := endpoints.UsersGetHandler(&ctx, request)
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "failed to retrieve users: failed to query user data"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("get users ok for one user", func(t *testing.T) {
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

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		actual := endpoints.UsersGetHandler(&ctx, request)
		expectedUsers := []models.User{
			{
				ID:              1,
				SpotifyUsername: "username1",
				LastfmUsername:  "username2",
				CreatedAt:       time,
				UpdatedAt:       time,
			},
		}
		expected := endpoints.APIResponse{200, "success", map[string][]models.User{"users": expectedUsers}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("get users no results", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())

		columns := []string{"id", "spotify_username", "lastfm_username", "created_at", "updated_at"}
		mockDB.ExpectQuery(
			"select * from users",
		).WillReturnRows(
			mockDB.NewRows(columns),
		)

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		actual := endpoints.UsersGetHandler(&ctx, request)
		expected := endpoints.APIResponse{200, "success", map[string][]models.User{"users": {}}}
		assert.Equal(t, &expected, actual)
	})
}

func TestGetUserByIDHandler(t *testing.T) {
	t.Run("get user fails with 404", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectQuery("select * from users where id=1").WillReturnError(pgx.ErrNoRows)

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "1"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expected := endpoints.APIResponse{404, "failed", map[string]string{"error": "no matching user in database"}}
		assert.Equal(t, &expected, actual)
	})

	t.Run("get user fails with 400", func(t *testing.T) {
		mockDB, err := pgxmock.NewConn(
			pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual),
		)
		assert.NoError(t, err)
		defer mockDB.Close(context.Background())
		mockDB.ExpectQuery("select * from users where id=1").WillReturnError(pgx.ErrNoRows)

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/abc", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "abc"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expected := endpoints.APIResponse{400, "failed", map[string]string{"error": "invalid user id"}}
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

		r, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "1"})

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

		columns := []string{"id", "spotify_username", "lastfm_username", "created_at", "updated_at"}
		time := time.Now().UTC()
		mockDB.ExpectQuery(
			"select * from users where id=1",
		).WillReturnRows(
			mockDB.NewRows(columns).AddRow(int64(1), "username1", "username2", time, time),
		)

		ctx := config.AppContext{
			AppConfig:    &config.AppConfig{},
			Context:      context.Background(),
			DBConnection: mockDB,
		}

		r, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		r = mux.SetURLVars(r, map[string]string{"id": "1"})

		actual := endpoints.UsersGetByIDHandler(&ctx, r)
		expectedUser := models.User{
			ID:              1,
			SpotifyUsername: "username1",
			LastfmUsername:  "username2",
			CreatedAt:       time,
			UpdatedAt:       time,
		}
		expected := endpoints.APIResponse{200, "success", map[string]models.User{"user": expectedUser}}
		assert.Equal(t, &expected, actual)
	})
}