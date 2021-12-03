package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Will return one router with routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
