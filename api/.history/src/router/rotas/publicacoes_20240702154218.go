package rotas

import "net/http"

var rotasPublicacoes = []Rota{
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
	}
}