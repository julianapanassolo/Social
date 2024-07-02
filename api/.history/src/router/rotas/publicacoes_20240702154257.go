package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	
}