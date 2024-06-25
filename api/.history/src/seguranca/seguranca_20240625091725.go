package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash = Irá receber uma string e coloca um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}