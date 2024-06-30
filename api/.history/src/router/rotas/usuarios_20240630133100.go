package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Rota representa uma estrutura de uma rota na API
type Rota struct {
    URI                string
    Metodo             string
    Funcao             func(http.ResponseWriter, *http.Request)
    RequerAutenticacao bool
}

// Definindo as rotas dos usuários
var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "usuarios/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	}
	{
		URI: "usuarios/{usuarioId}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},


