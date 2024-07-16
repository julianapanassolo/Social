package modelos

// DadosAutenticacao = contém o token e o id do usuário autenticado
type DadosAutenticacao struct {
	ID json. `json:"id"`
	Token string `json:"token"`
}