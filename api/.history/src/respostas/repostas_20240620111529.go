package respostas

import (
	"encoding/json"
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
	fmt.pr

}