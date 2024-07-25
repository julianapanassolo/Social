package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie


// Configurar = Utiliza as varáveis de ambiente para criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)	
}

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
		Name: "dados",
		Value: dadosCodificados,
		Path: "/",
		HttpOnly: true,
	})

	return nil 
}