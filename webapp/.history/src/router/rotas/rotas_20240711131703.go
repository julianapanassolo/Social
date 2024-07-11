package rotas

import "net/http"


// Rota = Representa todas as rotas 
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}