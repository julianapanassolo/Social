package router

import "github.com/gorilla/mux"


// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}