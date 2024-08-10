package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout = Remove os dados de autenticação salvos no browser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar()
}