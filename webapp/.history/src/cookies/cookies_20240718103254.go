package cookies

import (
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"

	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie


// Configurar = Utiliza as varáveis de ambiente para criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)	
}

// Salvar = Registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string {
		"id": ID,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dadosCookie", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name: "dadosCookie",
		Value: dadosCodificados,
		Path: "/",
		HttpOnly: true,
	})

	return nil 
}

func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dadosCookie")
	
}

