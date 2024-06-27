package middlewares

import (
	"fmt"
	"net/http"
)

func Autenticar(next http.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
	}
}