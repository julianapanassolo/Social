package respostas

import (
	"encoding/json"	
	"log"
	"net/http"
)

// JSON: Retorna uma reposta em JSON, para a requisição - Função de resposta ge
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro: Retorna um erro em formato JSON especifico
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		JSON(w, statusCode, struct {
			Erro string `json:"erro"`
		}){
			Erro: erro.Error
		}
	})

}