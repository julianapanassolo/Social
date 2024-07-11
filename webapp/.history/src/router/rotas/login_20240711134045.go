package rotas

import "net/http"

var rotasLogin = []Rota{
	{
		URI: "/",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	}
}