package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Rota de login

var rotaLogin = Rota{
	URI: "/login",
	Metodo: http.MethodPost,
	Funcao: controllers
	RequerAutenticacao: false,
}