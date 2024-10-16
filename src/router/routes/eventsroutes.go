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
	{
		Uri:      "/events/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetEventsByID,
	},
	{
		Uri:      "/events/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateEventByID,
	},
}
