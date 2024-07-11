package router

import "github.com/gorilla/mux"


// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")