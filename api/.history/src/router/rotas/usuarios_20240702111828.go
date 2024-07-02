package rotas

import (
	"api/src/controllers"
	"api/src/middlewares"
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
	{
		URI: "usuarios/{usuarioId}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "usuarios/{usuarioId}/seguidores",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI: "usuarios/{usuarioId}/seguindo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI: "usuarios/{usuarioId}/atualizar-senha",
		Metodo: http.MethodPost,
		Funcao: controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},

}

// Rota de login
var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	for _, rota := range rotasUsuarios {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	r.HandleFunc(rotaLogin.URI, middlewares.Logger(rotaLogin.Funcao)).Methods(rotaLogin.Metodo)

	return r
}

