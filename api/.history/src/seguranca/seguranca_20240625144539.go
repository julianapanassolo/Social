package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash = Irá receber uma string e coloca um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha = Vai comparar uma senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaSt string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}