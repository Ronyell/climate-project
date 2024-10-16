package routes

import (
	"api/src/controllers"
	"net/http"
)

var eventsRoutes = []Route{
	{
		Uri:      "/events",
		Method:   http.MethodPost,
		Function: controllers.CreateEvents,
	},
	{
		Uri:      "/events",
		Method:   http.MethodGet,
		Function: controllers.GetAllEvents,
	},
}
