package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Function responsible to create city
func CreateCities(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cityObj models.City
	if erro = json.Unmarshal(requestBody, &cityObj); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	if erro = validate.Struct(cityObj); erro != nil {
		errors := erro.(validator.ValidationErrors)
		fmt.Println(errors)
		response.Erro(w, http.StatusBadRequest, errors)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetCitiesRepository(db)

	cityObj.ID, erro = repository.Create(cityObj)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, cityObj)

}

// Function responsible to get all cities
func GetAllCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all cities"))
}

// Function responsible to get all cities from a UF
func GetCitiesByUF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get cities by UF"))
}
