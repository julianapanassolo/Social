package modelos

import (
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
)

// Usuário = Representa uma pessoa utilizando a aplicação
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json: "nick"`
	CriadoEm    time.Time    `json: "criadoEm"`
	Seguidores  []Usuario    `json: "seguidores"`
	Seguindo    []Usuario    `json: "seguindo"`
	Publicacoes []Publicacao `json: "publicacoes"`
}

// BuscarUsuarioCompleto = Faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)
}

func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {
	url: fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := 

}

func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {
	
}

func BuscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
