package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func CreateCitiesBulk(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cityObjts []models.City
	if erro = json.Unmarshal(requestBody, &cityObjts); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct

	for _, obj := range cityObjts {
		if erro = validate.Struct(obj); erro != nil {
			errors := erro.(validator.ValidationErrors)
			fmt.Println(errors)
			response.Erro(w, http.StatusBadRequest, errors)
			return
		}
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetCitiesRepository(db)

	for _, obj := range cityObjts {
		objId, erro := repository.Create(obj)
		obj.ID = fmt.Sprint(objId)
		if erro != nil {
			response.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}

	response.JSON(w, http.StatusCreated, cityObjts)

}

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

	objId, erro := repository.Create(cityObj)
	cityObj.ID = fmt.Sprint(objId)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, cityObj)

}

// Function responsible to get all cities
func GetAllCities(w http.ResponseWriter, r *http.Request) {
	ufSearch := r.URL.Query().Get("uf")

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetCitiesRepository(db)

	cities, erro := repository.GetCityByUF(ufSearch)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, cities)
}

// Function responsible to get city by ID
func GetCitiesByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	cityID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetCitiesRepository(db)

	cityObj, erro := repository.GetCityById(cityID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if cityObj.ID == "" {
		response.Erro(w, http.StatusNotFound, errors.New("city not found"))
		return
	}

	response.JSON(w, http.StatusOK, cityObj)
}

// Function responsible to update all city by id
func UpdateCityByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	cityID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

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

	if erro = repository.UpdateCityById(cityID, cityObj); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
