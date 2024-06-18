package router

import "github.com/gorilla/mux"


// "Gerar" = Irá retornar um router com as rotas configurações
func Gerar() *mux.Router {
	return mux.NewRouter()
}