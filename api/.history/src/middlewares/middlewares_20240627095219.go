package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger = Escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.pri("\n %s %s %s, r.Method, r.RequestURI, r.Host")
		next(w, r)
	}
}

// Atenticar = Verifica se o usuário no momento da requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando usuário...")
		next(w, r)
	}
}