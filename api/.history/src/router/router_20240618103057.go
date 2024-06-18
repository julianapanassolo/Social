package router

import "github.com/gorilla/mux"


// "Gerar" = Irá retornar
func Gerar() *mux.Router {
	return mux.NewRouter()
}