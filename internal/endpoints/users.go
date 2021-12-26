package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/models"
	"github.com/petermnhull/go-music-app/internal/repositories"
)

// UsersCountHandler returns number of users
func UsersCountHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	count, err := repositories.CountUsers(ctx.Context, ctx.DBConnection)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to get count of users")
	}
	data := map[string]string{
		"count": fmt.Sprintf("%v", count),
	}
	return NewAPIResponseSuccess(http.StatusOK, data)
}

// UsersGetHandler returns a list of current users
func UsersGetHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	users, err := repositories.GetAllUsers(ctx.Context, ctx.DBConnection)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to retrieve users: "+err.Error())
	}
	data := map[string][]models.User{
		"users": users,
	}
	return NewAPIResponseSuccess(http.StatusOK, data)
}

// UsersGetByIDHandler returns user matching ID
func UsersGetByIDHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	parameters := mux.Vars(r)
	idString := parameters["id"]
	if idString == "" {
		return NewAPIResponseFailed(http.StatusBadRequest, "no user id provided")
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return NewAPIResponseFailed(http.StatusBadRequest, "invalid user id")
	}

	user, err := repositories.GetUserByID(ctx.Context, ctx.DBConnection, id)
	if err != nil {
		var e *repositories.ErrNoRecords
		if errors.As(err, &e) {
			return NewAPIResponseFailed(http.StatusNotFound, e.Message)
		}
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to retrieve users: "+err.Error())
	}

	data := map[string]models.User{
		"user": *user,
	}
	return NewAPIResponseSuccess(http.StatusOK, data)
}
