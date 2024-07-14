package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CriarUsuario = Chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	nome := r.FormValue("nome")
	fmt.Println(nome)

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"nome": r.FormValue("nome"),
		"nome": r.FormValue("nome"),
		"nome": r.FormValue("nome"),

	})
}