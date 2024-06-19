package repositorios

import (
	"database/sql"
)

type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}


// NovoRepositorioDeUsuarios = Cria um repositório de usuários 
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}