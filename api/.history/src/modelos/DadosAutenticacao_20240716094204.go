package modelos


// DadosAutenticacao = contém o token e o id 
type DadosAutenticacao struct {
	ID string `json:"id"`
	Token string `json:"token"`
}