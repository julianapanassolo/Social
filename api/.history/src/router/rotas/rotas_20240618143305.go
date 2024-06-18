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

func Configurar(r *mux.Route) *mux.Route