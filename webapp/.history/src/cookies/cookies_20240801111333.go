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

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dadosCookie")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("dadosCookie", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil
}

