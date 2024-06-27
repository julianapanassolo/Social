package middlewares

import (
	"fmt"
	"net/http"
)

// Atenticar = Verifica 
func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(w, r)
	}
}