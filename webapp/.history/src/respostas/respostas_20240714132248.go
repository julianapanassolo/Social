package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI = Representa uma resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON = Retorna um resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode()

}