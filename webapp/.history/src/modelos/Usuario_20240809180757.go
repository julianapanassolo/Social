package modelos


// Usuário = Representa uma pessoa utilizando a aplicação
type Usuario struct {
	ID uint64 `json:"id`
	Nome string `json:"nome`
	Email string `json:"email"`
	Nick string `json: "nick`
}