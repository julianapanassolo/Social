package router

import (
	
	"github.com/gorilla/mux"
	"webapp/src/controllers" 
)

// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.CarregarTelaDeLogin).Methods("GET")
	r.HandleFunc("/", controllers.CarregarTelaDeLogin).Methods("GET") 	
	
	return r
}
