package router

import "github.com/gorilla/mux"


// Gerar = Reto
func Gerar() *mux.Router {
	return mux.NewRouter()
}