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
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}