package respostas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if err = json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func Erro() {
	if err = json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}

}