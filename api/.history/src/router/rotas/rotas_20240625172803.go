package rotas

import (
	"net/http"
	"github.com/gorilla/mux"
)

// "Rota" = Todas as rotas da API
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, * http.Request)
	RequerAutenticacao bool
}

// Configurar = Coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {

		}else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}
		
	}

	return r
}