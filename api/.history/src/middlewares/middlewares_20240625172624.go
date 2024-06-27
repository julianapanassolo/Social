package middlewares

import (
	"fmt"
	"net/http"
)

// Atenticar = Verifica se o usuário no momento da requisição está autenticado
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}