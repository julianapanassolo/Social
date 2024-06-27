package autenticacao

import (
	"time"
	"api/src/config"
	"fmt"
	"net/http"
	"strings"	
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken = Retorna tocken assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Token expira após 6 meses
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SECRET_KEY))
}

// ValidarToken = Verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	if tokenString == "" {
		return fmt.Errorf("Token não encontrado ou mal formatado")
	}
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Cl

	fmt.Println(token)
	return nil
}


func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	fmt.Printf("Authorization header: %s\n", token) // Log para verificar o valor do cabeçalho
	if token == "" {
		return ""
	}
	
	// Verificação adicional de caracteres inválidos
	if strings.ContainsAny(token, "\n\r") {
		fmt.Println("O cabeçalho de autorização contém caracteres inválidos")
		return ""
	}

	// Verifica o formato do token
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		fmt.Println("O formato do cabeçalho de autorização deve ser Bearer {token}")
		return ""
	}

	return parts[1]
}


func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SECRET_KEY, nil

}