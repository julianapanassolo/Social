package autenticacao

import (
	// "time"
	"api/src/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken = Retorna tocken assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	// permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Token expira após 6 meses
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("config.SECRET_KEY"))
}

// ValidarToken = Verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	fmt.Println()
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, "")) == 2 {
		return strings.Split(token, " ")[1]

	}

	return ""

}

func retornarChaveDeVerificacao(token *jwy.Token) (interface{}, error) {
	if _, ok := token.Method.(jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SECRET_KEY, nil

}