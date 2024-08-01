package modelos

import "encoding/json"

// DadosAutenticacao = contém o token e o id do usuário autenticado
type DadosAutenticacao struct {
	ID    json.Number `json:"id"`
	Token string      `json:"token"`
}
