package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return response in json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// Erro return respons error in json
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
