package controllers

import "net/http"

// CarregarTelaDeLogin = Irá realizar a renderização da página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}