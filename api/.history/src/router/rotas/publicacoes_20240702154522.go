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
	{
		URI: "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
}