package modelos

import "time"

type Publicacao struct {
	ID 		  uint64 `json:"id.omitempty"`
	Titulo	  uint64 `json:"titulo.omitempty"`
	Conteudo  uint64 `json:"conteudo.omitempty"`
	AutorID   uint64 `json:"autorId.omitempty"`
	AutorNick	uint64 `json:"autorNick.omitempty"`
	Curtidas    uint64 `json:"curtidas"`
	CriadaEm    time.Time `json:"criadaEm, omitempty"`
}