package rotas

import (
	"api/src/controllers"
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
// 	{
// 		URI: "/publicacoes",
// 		Metodo: http.MethodGet,
// 		Funcao: controllers.BuscarPublicacao,},
// 		RequerAutenticacao: true,
// 	},
// 	{
// 		URI: "/publicacoes/{publicacaoId}",
// 		Metodo: http.MethodGet,
// 		Funcao: controllers.BuscarPublicacoes,
// 		RequerAutenticacao: true,
// 	},
// 	{
// 		URI: "/publicacoes/{publicacaoId}",
// 		Metodo: http.MethodPut,
// 		Funcao: controllers.AtualizarPublicacao,
// 		RequerAutenticacao: true,
// 	},
// 	{
// 		URI: "/publicacoes/{publicacaoId}",
// 		Metodo: http.MethodDelete,
// 		Funcao: controllers.DeletarPublicacao,
// 		RequerAutenticacao: true,
// 	},
}