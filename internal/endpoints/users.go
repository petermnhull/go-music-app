package endpoints

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/internal/repositories"
)

// UsersGetByIDHandler returns user matching ID
func UsersGetByIDHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	// Get request parameters
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		return NewAPIResponseFailed(http.StatusBadRequest, "no id provided")
	}

	// Get user from database
	user, err := repositories.GetUserByID(ctx.Context, ctx.DBConnection, id)
	if err != nil {
		var e *repositories.ErrNoRecords
		if errors.As(err, &e) {
			return NewAPIResponseFailed(http.StatusNotFound, e.Message)
		}
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to retrieve users: "+err.Error())
	}

	// Return user data
	data := map[string]models.User{
		"user": *user,
	}
	return NewAPIResponseSuccess(http.StatusOK, data)
}
