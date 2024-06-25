package modelos

import (
	"errors"
	"strings"
	"time"
	"github.com/badoux/checkmail"
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

// Preparar = Irá chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	usuario.formatar(etapa)
	return nil 
	
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo NOME é obrigatório!")
	}	
	if usuario.Nick == "" {
		return errors.New("O campo NICK é obrigatório!")
	}
	if usuario.Email == "" {
		return errors.New("O campo E-MAIL é obrigatório!")
	}
	
	if erro:= checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo SENHA é obrigatório!")
	}

	if etapa == "cadastro" && len(usuario.Senha) > 6 {
		return errors.New("Senha deve ser no máximo 6 digitos!")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro"
}

