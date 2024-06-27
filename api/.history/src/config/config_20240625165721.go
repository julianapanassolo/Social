package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = ""
	// Porta é onde a API vai estar rodando
	Porta = 0

	// SECRET_KEY = É a chave que vai ser usada para assinar o token
	SECRET_KEY []byte
)

// Carregar = Vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	// Carregar as variáveis de ambiente do arquivo .env
	if err = godotenv.Load(); err != nil {
		log.Fatal(err) // Se houver problema em carregar as variáveis de ambiente, a API não roda
	}

	// Converter a string da variável de ambiente API_PORT para um número inteiro
	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000 // Porta padrão se a variável de ambiente não estiver definida ou for inválida
	}

	// Linha que cria a string de conexão com o banco
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),

		SECRET_KEY
	)
}
