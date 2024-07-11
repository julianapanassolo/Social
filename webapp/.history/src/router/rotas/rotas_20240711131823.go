package rotas

import (
	"net/http"
	"webapp/src/router"

	"github.com/gorilla/mux"
)

// Rota = Representa todas as rotas da Aplicação Web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// C
func Configurar(router *mux.Router) *mux.Router {

}