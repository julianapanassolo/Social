package modelos

import (
	"errors"
	"strings"
	"time"
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

func (usuario *Usuario) preparar() {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	
	
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo NOME é obrigatório!")
	}	
	if usuario.Nick == "" {
		return errors.New("O campo NICK é obrigatório!")
	}
	if usuario.Email == "" {
		return errors.New("O campo E-MAIL é obrigatório!")
	}
		if usuario.Senha == "" {
		return errors.New("O campo SENHA é obrigatório!")
	}
	if len(usuario.Senha) > 6 {
		return errors.New("Senha deve ser no máximo 6 digitos!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

