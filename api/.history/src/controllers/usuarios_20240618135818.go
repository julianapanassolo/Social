package controllers

import "net/http"

// CriarUsuario = Insere 1 usuário na banco de dados
func CriarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Criando Usuário"))
}
// BuscarUsuarios = Busca 
func BuscarUsuarios (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando todos os Usuário"))
}

func BuscarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando um usuário!"))
}

func AtualizarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Atualizando Usuário"))
}
func DeletarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Deletando Usuário"))
}