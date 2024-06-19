package config

import "github.com/joho/godotenv"

var (
	// StringConexaoBanco = É a string de conexeção com o mysql
	StringConexaoBanco = ""
	// Porta: Onde a API vai estar rodando
	Porta = 0 
)


// Carregar: Vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if  erro = godotenv


}