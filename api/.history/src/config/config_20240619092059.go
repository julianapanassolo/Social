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
	var erro error

	if  erro = godotenv.Load(); erro != nil {
		log.Fatal(erro) // Se houver problema em carregar as variávies de ambiente, a api não roda
	}

	Porta, erro = strconv.Atoi(os.Getenv(("API_PORT"))) // converte uma string para número inteiro
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintln("%s:%s@/%s?charset=utf8parseTime=True&loc=Local",
	os.Getenv("DB_USUARIO"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"),

)

}