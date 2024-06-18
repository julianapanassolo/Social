package router

import "github.com/gorilla/mux"


// "Gerar" = Irá retornar um router com as configurações
func Gerar() *mux.Router {
	return mux.NewRouter()
}