package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("\n %s %s %s, r")
	}
}

// Atenticar = Verifica se o usuário no momento da requisição está autenticado
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando usuário...")
		next(w, r)
	}
}