package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"webapp/src/controllers" // Certifique-se de ter o controlador importado corretamente
)

// Gerar = Retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.CarregarTelaDeLogin).Methods("GET") // Adiciona uma rota para a página inicial
	// Adicione outras rotas conforme necessário
	return r
}
