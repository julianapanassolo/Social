package controllers

import "net/http"

// CriarUsuario = Insere 1 usuário na banco de dados
func CriarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Criando Usuário"))
}
// BuscarUsuarios = Busca todos os usuários salvos no banco de dados
func BuscarUsuarios (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando todos os Usuário"))
}
// BuscarUsuario = Busca apenas 1 usuário salvo no banco de dados
func BuscarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando um usuário!"))
}
// AtualizarUsuario = Altera as informações de um usuário no banco
func AtualizarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Atualizando Usuário"))
}
// DeletarUsuario = Excluiu as informações de um usuário no banco de dados
func DeletarUsuario (w http.ResponseWriter r *http) {
	w.Write([]byte("Deletando Usuário"))
}