package middlewares

import (
	"fmt"
	"net/http"
)

// Atenticar = 
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(w, r)
	}
}