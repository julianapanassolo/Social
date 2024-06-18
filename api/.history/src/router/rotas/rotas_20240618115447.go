package rotas

import "net/http"

type Rota struct {
	Uri string
	Metodo string
	Funcao func(http.ResponseWriter, * http.Request)
	Requer
}