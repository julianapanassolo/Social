package modelos

type Publicacao struct {
	ID uint64 `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty"`
	Conteudo string `json:"conteudo,omitempty"`
	AutorID uint64 `json:"autorId,omitempty"`
	Autor uint64 `json:"id,omitempty"`

	
}