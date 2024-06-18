package rotas

type Rota struct {
	Uri string
	Metodo string
	Funcao func(http.Res)
}