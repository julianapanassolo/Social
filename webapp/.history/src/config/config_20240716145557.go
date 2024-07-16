package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// APIURL - Representa a URL para comunicação com a API
	APIURL = ""

	// Porta = Onde a aplicação web está rodando
	Porta = 0

	// HashKey = É utilizada para autenticar o cookie
	HashKey []byte

	// BlockKey = É utilizada para ciptografar os dados do cookies 
	BlockKey []byte
)

// Carregar = Inicializa as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte()

}