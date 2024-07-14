package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

var rotasUsuarios = []Rota {
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
