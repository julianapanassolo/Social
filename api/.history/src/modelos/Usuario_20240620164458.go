package modelos

import (
	"time"
	"errors"
)

// Usuario = Representa um usuário utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json: "nome,omitempty"`
	Nick string `json: "nick,omitempty"`
	Email string `json: "email,omitempty"`
	Senha string `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`

}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo NOME é obrigatório!)
	}	
	if usuario.Nick == "" {
		return errors.New("O campo NICK é obrigatório!)
	}
	if usuario.Email == "" {
		return errors.New("O campo E-MAIL é obrigatório!)
	}
		if usuario.Senha == "" {
		return errors.New("O campo SENHA é obrigatório!)
	}
	if usuario.Senha == "{=< 6}" {
		return errors.New("Senha de no máximo 6 digitos!)
	}
}

