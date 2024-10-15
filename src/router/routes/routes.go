package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represent all routes in API
type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

// Configure all routes
func Configure(r *mux.Router) *mux.Router {
	routes := citiesRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
