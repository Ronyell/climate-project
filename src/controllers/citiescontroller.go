package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Function responsible to create city
func CreateCities(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var city models.City
	if erro = json.Unmarshal(requestBody, &city); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.GetCitiesRepository(db)

	cityID, erro := repository.Create(city)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Create city with id %d", cityID)))

}

// Function responsible to get all cities
func GetAllCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all cities"))
}

// Function responsible to get all cities from a UF
func GetCitiesByUF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get cities by UF"))
}
