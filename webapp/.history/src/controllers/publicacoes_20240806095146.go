package controllers

import (
	"fmt"
	"net/http"
)

// CriarPublicacao = Chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	publicacao, erro := json.
}