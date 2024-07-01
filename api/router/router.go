package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
