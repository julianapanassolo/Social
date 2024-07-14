package router

import (
	
	"github.com/gorilla/mux"
	"webapp/src/controllers" 
)

// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}).Methods("GET")

	r.HandleFunc("/login", controllers.CarregarTelaDeLogin).Methods("GET")
	
	return r
}
