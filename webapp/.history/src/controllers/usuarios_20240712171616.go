package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// CriarUsuario = Chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	fmt.Print(usuario)
}