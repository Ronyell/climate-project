package controllers

import "net/http"

// Function responsible to create city
func CreateCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create city"))
}

// Function responsible to get all cities
func GetAllCities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all cities"))
}

// Function responsible to get all cities from a UF
func GetCitiesByUF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get cities by UF"))
}
