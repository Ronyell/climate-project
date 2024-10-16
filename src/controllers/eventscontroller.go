package controllers

import (
	"api/src/database"
	"api/src/models"
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
	eventType := r.URL.Query().Get("type")
	cityUf := r.URL.Query().Get("uf")

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.GetEventsRepository(db)
	var events []models.EventDescriber

	if cityUf == "" {
		events, erro = repository.GetEventByType(eventType)
	} else {
		events, erro = repository.GetEventByTypeAndUf(eventType, cityUf)
	}

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, events)
}
