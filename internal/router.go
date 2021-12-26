package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
)

type appHandler struct {
	AppContext *config.AppContext
	Handler    func(ctx *config.AppContext, r *http.Request) *endpoints.APIResponse
}

// ServeHTTP to satisfy the http.Handler interface
func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := ah.Handler(ah.AppContext, r)
	w.WriteHeader(int(response.Code))
	output := response.ToOutput()
	fmt.Fprint(w, output)
}

// NewRouter provides handler for all endpoints
func NewRouter(ctx *config.AppContext) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/health", appHandler{ctx, endpoints.HealthCheckHandler}).Methods(http.MethodGet)
	r.Handle("/users", appHandler{ctx, endpoints.UsersGetHandler}).Methods(http.MethodGet)
	r.Handle("/users", appHandler{ctx, endpoints.UserUpsertHandler}).Methods(http.MethodPost)
	r.Handle("/users/{id:[0-9+]}", appHandler{ctx, endpoints.UsersGetByIDHandler}).Methods(http.MethodGet)
	r.Handle("/users/count", appHandler{ctx, endpoints.UsersCountHandler}).Methods(http.MethodGet)
	return r
}
