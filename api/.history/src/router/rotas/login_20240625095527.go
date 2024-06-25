package rotas

import "net/http"

// Rota de login

var rotaLogin = Rota{
	URI: "/login",
	Metodo: http.MethodPost,
	Funcao: contro
	RequerAutenticacao: false,
}