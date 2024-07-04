package modelos

import (
	"time"

	"golang.org/x/crypto/openpgp/errors"
)

// Publicacao = Representa uma publicação feita por um usuário
type Publicacao struct {
	ID 		    uint64 `json:"id.omitempty"`
	Titulo	    uint64 `json:"titulo.omitempty"`
	Conteudo    uint64 `json:"conteudo.omitempty"`
	AutorID     uint64 `json:"autorId.omitempty"`
	AutorNick	uint64 `json:"autorNick.omitempty"`
	Curtidas    uint64 `json:"curtidas"`
	CriadaEm    time.Time `json:"criadaEm, omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil 
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteúdo")
}