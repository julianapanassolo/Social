package router

import (
	"net/http"
	"webapp/src/controllers"

	"github.com/gorilla/mux"
)

// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}).Methods("GET")

	r.HandleFunc("/login", controllers.CarregarTelaDeLogin).Methods("GET")

	fs := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	
	return r
}
