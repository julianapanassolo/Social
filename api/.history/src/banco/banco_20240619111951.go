package banco

import (
	"api/src/con"
	"database/sql"
	_"github.com/go-sql-driver/mysql" // Driver
)

// Conectar = Abre a conexão com o Banco de Dados e retorna
func Conectar() (*sql.DB, error) {
	dbm erro := sql.Open("mysql", config.StringConexaoBanco)


}