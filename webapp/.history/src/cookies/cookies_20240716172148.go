package cookies

import (
	"webapp/src/config"
	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie


// Configurar = Utiliza 
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)	
}