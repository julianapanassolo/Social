package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc)

// Atenticar = Verifica se o usuário no momento da requisição está autenticado
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando usuário...")
		next(w, r)
	}
}