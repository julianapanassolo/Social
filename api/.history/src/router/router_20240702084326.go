package router

import (
	"api/src/controllers"
	"api/src/middlewares"
	"github.com/gorilla/mux"
)

// Gerar = Irá retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/usuarios", controllers.CriarUsuario).Methods("POST")

	return r
}


