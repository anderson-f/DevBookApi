package router

import "github.com/gorilla/mux"

// Will return one router with routes configured
func Generate() *mux.Router {
	return mux.NewRouter()
}
