package rotas

import "net/http"

var rotasUsuarios = []Rota {
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
	}
}