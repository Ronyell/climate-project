package controllers

import (
	"api/src/database"
	"api/src/repositories"
	"api/src/response"
	"io"
	"net/http"
)

// Function responsible to create city
func CreateEvents(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetEventsRepository(db)

	_, erro = repository.CreateEvent(requestBody)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Function responsible to get all cities
func GetAllEvents(w http.ResponseWriter, r *http.Request) {

}

// Function responsible to get city by ID
func GetEventsByID(w http.ResponseWriter, r *http.Request) {

}

// Function responsible to update all city by id
func UpdateEventByID(w http.ResponseWriter, r *http.Request) {

}
