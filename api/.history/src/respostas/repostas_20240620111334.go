package respostas

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if erro = json.NewEncoder(w).Encode(dados); erro != nil {
		
	}
}

func Erro() {

}