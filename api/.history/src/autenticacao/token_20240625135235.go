package autenticacao

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	
}