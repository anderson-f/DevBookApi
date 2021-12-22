package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

// Configure put all routes inside router
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routeLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
