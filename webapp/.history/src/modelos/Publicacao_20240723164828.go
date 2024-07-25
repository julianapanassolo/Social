package modelos

type Publicacao struct {
	ID uint64 `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty"`
	Conteudo string `json:"conteudo,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	AutorID uint64 `json:"autorId,omitempty"`
	Curtidas uint64 `json:"curtidas"`

	
}