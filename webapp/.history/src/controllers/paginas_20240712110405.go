package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin = Irá realizar a renderização da página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
	
}

// CarregarPaginaDeCadastroDeUsuario = Irá cadastrar
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {

}