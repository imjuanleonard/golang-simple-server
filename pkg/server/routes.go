package server

import (
	"github.com/gorilla/mux"
	"github.com/imjuanleonard/palu-covid/pkg/middleware"
	h "github.com/imjuanleonard/palu-covid/internal/handler"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.SkipClean(true)
	r.Use(middleware.Recover())
	r.HandleFunc("/ping", h.PingHandler).Methods(http.MethodGet)
	return r
}

func NewRouter(handler *handler) *mux.Router {
	r := newRouter()

	// TODO: Put all router HERE
	return r
}
