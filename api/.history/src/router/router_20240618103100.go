package router

import "github.com/gorilla/mux"


// "Gerar" = Irá retornar um router
func Gerar() *mux.Router {
	return mux.NewRouter()
}