package repositorios

import (
	"database/sql"
)

// Usuarios: Representa um repositório de usuários
type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}


// NovoRepositorioDeUsuarios = Cria um repositório de usuários 
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}	
}

// Criar = Insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
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

	ultimoIdInserido, erro := resultado.LastIn
}