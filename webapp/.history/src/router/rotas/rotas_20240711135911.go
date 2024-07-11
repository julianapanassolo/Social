package rotas

import (
	"net/http"
	

	"github.com/gorilla/mux"
)

// Rota = Representa todas as rotas da Aplicação Web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar = Coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	for _, rota := rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota)
	}

}