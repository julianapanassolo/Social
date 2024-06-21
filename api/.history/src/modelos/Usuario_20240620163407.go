package modelos

import (
	"time"
)

// Usuario = Representa um usuário utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json: "nome, omitempty"`
	Nick string `json: "nick,omitempty"`
	Email string `json: "email, omitempty"`
	Senha string `json: "senha, omitempty"`
	CriadoEm time.Time `json: "CriadoEm, omitempty"`

}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo ")
	}
}