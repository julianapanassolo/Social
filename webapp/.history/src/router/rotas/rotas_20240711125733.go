package rotas

type Rota struct {
	URI string
	Metodo string
	Funcao func(h)
}