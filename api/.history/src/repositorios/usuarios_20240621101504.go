package repositorios

import (
	"database/sql"
	"api/src/modelos"
	"fmt"
)

// Usuarios: Representa um repositório de usuários
type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}


// NovoRepositorioDeUsuarios = Cria um repositório de usuários 
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db: db}	
}

// Criar = Insere um usuário no banco de dados
func (repositorio *usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",	
	)
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar = Irá trazer todos os usuários que atendem um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, CriadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", 
		fmt.Sprintf("%%%s%%", nomeOuNick), 
		fmt.Sprintf("%%%s%%", nomeOuNick),
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorio *)
