package internal

import (
	"net/http"

	"github.com/petermnhull/go-music-app/internal/config"
)

// AppHandler wraps endpoint handlers with context
type AppHandler struct {
	AppContext *config.AppContext
	Handler    func(ctx *config.AppContext, w http.ResponseWriter, r *http.Request)
}

func newHandler(ctx *config.AppContext, endpointHandler func(ctx *config.AppContext, w http.ResponseWriter, r *http.Request)) AppHandler {
	return AppHandler{ctx, endpointHandler}
}

// ServeHTTP to satisfy the http.Handler interface
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.Handler(ah.AppContext, w, r)
}
