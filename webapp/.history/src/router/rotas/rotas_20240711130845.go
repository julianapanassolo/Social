package rotas

import "net/http"

type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAtenticacao bool

} "go.formatTool": "goimports",
  "editor.formatOnSave": true "go.formatTool": "goimports",
  "editor.formatOnSave": true