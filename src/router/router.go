package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate all routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
