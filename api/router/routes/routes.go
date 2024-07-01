package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path         string
	Method       string
	Function     func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func Configure(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.HandleFunc(route.Path, route.Function).Methods(route.Method)
	}
	return r
}
