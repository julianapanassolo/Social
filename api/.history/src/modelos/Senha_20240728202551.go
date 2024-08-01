package modelos


// DadosAutenticacao contém o token e o id do usuário autenticado
type DadosAutenticacao struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}


// Senha = Representa o formato da requisição de alteração de senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}

