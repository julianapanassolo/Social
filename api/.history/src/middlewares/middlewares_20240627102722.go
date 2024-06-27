package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"fmt"
	"log"
	"net/http"
)

// Logger = Escreve informações da requisição no terminal
func Logger( http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		(w, r)
	}
}

// Atenticar = Verifica se o usuário no momento da requisição está autenticado
func Autenticar( http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		(w, r)
	}
}