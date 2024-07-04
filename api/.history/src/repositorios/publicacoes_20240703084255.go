package repositorios

import "database/sql"


// Publicacoes = Representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}


// NovoRepositorioDePublicacoes = Cria 
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}