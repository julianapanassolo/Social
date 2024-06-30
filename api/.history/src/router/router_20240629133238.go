package router

import (
	"api/src/controllers"
	"api/src/middlewares"
	"github.com/gorilla/mux"
)

// Gerar = Irá retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/usuario/{id}", middlewares.Autenticar(controllers.CriarUsuario)).Methods("GET")
	return r
}


// Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/login", controllers.Login).Methods("POST")

    r.HandleFunc("/usuario/{id}", middlewares.Autenticar(controllers.CriarUsuario)).Methods("GET")
    return r
}