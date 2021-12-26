package internal

import (
	"fmt"
	"net/http"

	"github.com/petermnhull/go-music-app/internal/config"
	"github.com/petermnhull/go-music-app/internal/endpoints"
)

// AppHandler wraps endpoint handlers with context
type AppHandler struct {
	AppContext *config.AppContext
	Handler    func(ctx *config.AppContext, r *http.Request) *endpoints.APIResponse
}

// ServeHTTP to satisfy the http.Handler interface
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := ah.Handler(ah.AppContext, r)
	w.WriteHeader(int(response.Code))
	output := response.ToOutput()
	fmt.Fprint(w, output)
}
