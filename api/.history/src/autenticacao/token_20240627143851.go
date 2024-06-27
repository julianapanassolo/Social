package autenticacao

import (
	// "time"
	"api/src/autenticacao"
	"api/src/respostas"
	"net/http"

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
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
		}
	}
	return nil
}