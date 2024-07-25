package rotas

import "net/http"

var rotaPaginaPrincipal = Rota{
	URI: "/home",
	Metodo: http.MethodGet,
}