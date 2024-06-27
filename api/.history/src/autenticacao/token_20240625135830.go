package autenticacao

import (
	// "time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	// permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Token expira após 6 meses
	permissoes["usuarioId"] = usuarioID
	
}