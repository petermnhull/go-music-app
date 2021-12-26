package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
)

func newRouter(ctx *config.AppContext) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/health", AppHandler{ctx, endpoints.HealthCheckHandler}).Methods(http.MethodGet)
	r.Handle("/users", AppHandler{ctx, endpoints.UsersGetHandler}).Methods(http.MethodGet)
	r.Handle("/users/{id:[0-9+]}", AppHandler{ctx, endpoints.UsersGetByIDHandler}).Methods(http.MethodGet)
	r.Handle("/users/count", AppHandler{ctx, endpoints.UsersCountHandler}).Methods(http.MethodGet)
	return r
}

// NewServer builds a new http.Server from context
func NewServer(ctx *config.AppContext) *http.Server {
	router := newRouter(ctx)
	address := fmt.Sprintf("127.0.0.1:%v", ctx.AppConfig.Port)
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: ctx.AppConfig.ReadTimeout,
		ReadTimeout:  ctx.AppConfig.WriteTimeout,
	}
	return server
}
