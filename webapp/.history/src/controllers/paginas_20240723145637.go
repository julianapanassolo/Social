package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin = Irá realizar a renderização da página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
	
}

// CarregarPaginaDeCadastroDeUsuario = Irá carregar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal = Carrega a página principal com as publicações - home
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintln()

}
