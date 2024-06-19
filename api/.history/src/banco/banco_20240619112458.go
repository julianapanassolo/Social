package banco

import (
	"api/src/config"
	"database/sql"
	_"github.com/go-sql-driver/mysql" // Driver
)

// Conectar = Abre a conexão com o Banco de Dados e retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro !


}