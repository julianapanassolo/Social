package middlewares

import (
	"log"
	"net/http"
)

//  Logger = Escreve informações da requisição
func Logger(proximaFuncao http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s, r.Method, r.RequestURI, r.Host")
		proximaFuncao(w, r)
	}
}