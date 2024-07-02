package router

import (
	"api/src/middlewares"
	"api/src/rotas"
	"github.com/gorilla/mux"
	"net/http"
)

// Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	// Configuração das rotas
	r = rotas.Configurar(r)

	return r
}


