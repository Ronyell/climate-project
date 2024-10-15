package routes

import (
	"api/src/controllers"
	"net/http"
)

var citiesRoutes = []Route{
	{
		Uri:      "/cities",
		Method:   http.MethodPost,
		Function: controllers.CreateCities,
	},
	{
		Uri:      "/cities",
		Method:   http.MethodGet,
		Function: controllers.GetAllCities,
	},
	{
		Uri:      "/cities/{uf}",
		Method:   http.MethodGet,
		Function: controllers.GetCitiesByUF,
	},
}
