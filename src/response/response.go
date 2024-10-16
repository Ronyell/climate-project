package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return response in json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		jsonData, erro := json.Marshal(data)
		if erro != nil {
			log.Fatal(erro)
		}
		w.Write(jsonData)
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
