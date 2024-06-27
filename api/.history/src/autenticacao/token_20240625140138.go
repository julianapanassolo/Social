package autenticacao

import (
	// "time"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken = Retorna tocken assinado
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	// permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Token expira após 6 meses
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString([]byte())
}