package middlewares

import (
	"fmt"
	"net/http"
)

// Atenticar = Verifica se o usuário no momento da requi
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(w, r)
	}
}