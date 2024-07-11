package rotas

import "net/http"
import "webapp/src/contro"

var rotasLogin = []Rota{
	{
		URI: "/",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,

	},
	{
		URI: "/login",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,

	},

}