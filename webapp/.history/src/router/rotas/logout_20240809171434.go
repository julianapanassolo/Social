package rotas

import "net/http"

var rotaLogout = Rota{
	URI: "/logout",
	Metodo: http.MethodGet,
	
}

