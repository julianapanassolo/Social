package modelos


// Senha = Representa o formato da requisição
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}