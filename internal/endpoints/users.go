package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/validator.v2"

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
	// Get request parameters
	params := mux.Vars(r)
	idString := params["id"]
	if idString == "" {
		return NewAPIResponseFailed(http.StatusBadRequest, "no user id provided")
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return NewAPIResponseFailed(http.StatusBadRequest, "invalid user id")
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

// UserUpsertHandler upserts a new user
func UserUpsertHandler(ctx *config.AppContext, r *http.Request) *APIResponse {
	// Get request body parameters
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to read request body")
	}

	// Decode body into expected parameters
	type upsertUserBody struct {
		SpotifyUsername string `json:"spotify_username" validate:"nonzero"`
		LastfmUsername  string `json:"lastfm_username,omitempty"`
	}
	var body upsertUserBody
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return NewAPIResponseFailed(http.StatusBadRequest, "invalid json request body")
	}

	// Validate parameters
	errs := validator.Validate(body)
	if errs != nil {
		return NewAPIResponseFailed(http.StatusBadRequest, "invalid parameters in request body")
	}

	// Upsert user
	user := models.User{SpotifyUsername: body.SpotifyUsername, LastfmUsername: body.LastfmUsername}
	err = repositories.UpsertUser(ctx.Context, ctx.DBConnection, &user)
	if err != nil {
		return NewAPIResponseFailed(http.StatusInternalServerError, "failed to upsert user: "+err.Error())

	}

	// Return data
	data := map[string]interface{}{
		"message":          "user upserted",
		"spotify_username": user.SpotifyUsername,
		"lastfm_username":  user.LastfmUsername,
	}
	return NewAPIResponseSuccess(http.StatusCreated, data)
}
