package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco = É a string de conexeção com o mysql
	StringConexaoBanco = ""
	// Porta: Onde a API vai estar rodando
	Porta = 0 
)

// Carregar: Vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	if  err = godotenv.Load(); err != nil {
		log.Fatal(err) // Se houver problema em carregar as variávies de ambiente, a api não roda
	}

	 // converte uma string para número inteiro
	if err != nil {
		Porta = 9000
	}

	// Linha que cria string de conexão com o banco
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8parseTime=True&loc=Local",
	os.Getenv("DB_USUARIO"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"),

)

}