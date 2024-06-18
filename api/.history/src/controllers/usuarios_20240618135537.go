package controllers

import "net/http"

func CriarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Criando Usuário"))
}

func BuscarUsuarios (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando todos os Usuário"))
}

func BuscarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Buscando um usuário!" Usuário"))
}

func CriarUsuario (w http.ResponseWriter r *http.Response) {
	w.Write([]byte("Criando Usuário"))
}