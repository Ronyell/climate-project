package routes

import (
	"api/src/middlewares"
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
	for _, route := range citiesRoutes {
		r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	for _, route2 := range eventsRoutes {
		r.HandleFunc(route2.Uri, middlewares.Logger(route2.Function)).Methods(route2.Method)
	}

	return r
}
