package router

import "github.com/gorilla/mux"


// Gerar = Retorna 
func Gerar() *mux.Router {
	return mux.NewRouter()
}