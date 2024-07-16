package cookies

import (
	"webapp/src/config"
	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie


// Configurar = Utiliza as varáveis de ambiente para criação do Se
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)	
}