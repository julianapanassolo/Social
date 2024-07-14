package controllers

import "net/http"


// CriarUsuario = Chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	nome := r.FormValue("nome")
	
}