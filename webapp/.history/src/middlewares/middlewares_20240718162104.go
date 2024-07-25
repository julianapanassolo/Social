package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/cookies"
)

//  Logger = Escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s, r.Method, r.RequestURI, r.Host")
		proximaFuncao(w, r)
	}
}

// Autenticar = Verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if_, erro := cookies.Ler(r); erro 

		
		proximaFuncao(w,r)

	}
}