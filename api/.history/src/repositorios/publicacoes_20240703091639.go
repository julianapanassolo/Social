package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes = Representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}


// NovoRepositorioDePublicacoes = Cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar = Insere uma nova publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement

}