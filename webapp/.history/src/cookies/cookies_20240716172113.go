package cookies

import (
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config)	
}