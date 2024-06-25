package controllers

import (
	"api/src/modelos"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login = Resposável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != {
		reps
	}

}