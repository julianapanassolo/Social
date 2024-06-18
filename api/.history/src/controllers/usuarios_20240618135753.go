package controllers

import "net/http"

// CriarUsu
func CriarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Criando Usuário"))
}

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