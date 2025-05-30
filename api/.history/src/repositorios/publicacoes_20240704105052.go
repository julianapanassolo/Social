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

	defer statement.Close()

	resultado, erro:= statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil 

}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p*, u.nick FROM
		publicacoes p INNER JOIN usuarios u
		ON u.id = p.autor_id WHERE p.id = ?,
		publicacoes
	
	`)
}