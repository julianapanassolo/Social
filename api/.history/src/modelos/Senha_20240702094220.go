package modelos


// Senha = Representa o formato da requisição de alteração de senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}

type DadosAutenticacao struct {
    ID    uint64 `json:"id"`
    Token string `json:"token"`
}
