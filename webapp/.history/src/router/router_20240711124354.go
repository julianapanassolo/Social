package router

import "github.com/gorilla/mux"


// Gerar = Retorna um router
func Gerar() *mux.Router {
	return mux.NewRouter()
}