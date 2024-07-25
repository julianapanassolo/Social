package middlewares

import (
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("")
	}
}