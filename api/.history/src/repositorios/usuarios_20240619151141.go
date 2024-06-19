package repositorios

import (
	"database"
)

type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}