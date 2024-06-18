package router

import "github.com/gorilla/mux"


// "Gerar" = Irá re
func Gerar() *mux.Router {
	return mux.NewRouter()
}