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
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Println("Erro ao codificar dados JSON:", erro)
		}
	}
	
}

// Erro: Retorna um erro especifico em formato JSON 
func Erro(w http.ResponseWriter, statusCode int, erro error) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    if erro != nil {     // Verifica se o erro não é nulo antes de acessar erro.Error()
        if err := json.NewEncoder(w).Encode(struct {
            Erro string `json:"erro"`
        }{
            Erro: erro.Error(),
        }); err != nil {
            // Se ocorrer um erro ao codificar o JSON, pode ser útil logar o erro ou retorná-lo, mas não use log.Fatal() aqui
            log.Println(err)
        }
    }
}
