package respostas

import (
	"encoding/json"	
	"log"
	"net/http"
)

// JSON: Retorna uma reposta em JSON para a requisição - Função de resposta generica
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {

	}


	
}

// Erro: Retorna um erro especifico em formato JSON 
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {		
			Erro string `json:"erro"`
		}{
			Erro: erro.Error(),
		})
	}