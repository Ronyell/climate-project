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
		Uri:      "/cities/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetCitiesByID,
	},
	{
		Uri:      "/cities/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateCityByID,
	},
	{
		Uri:      "/cities/bulk",
		Method:   http.MethodPost,
		Function: controllers.CreateCitiesBulk,
	},
}
