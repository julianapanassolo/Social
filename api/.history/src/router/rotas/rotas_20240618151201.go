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

// Configurar = 
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}