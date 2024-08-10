package rotas

import (
	"net/http"
	"webapp/src/con"
)

var rotaLogout = Rota{
	URI:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}
