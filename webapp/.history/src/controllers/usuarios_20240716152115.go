package controllers

import (
	"bytes"
	"encoding/json"	
	"net/http"
	"webapp/src/respostas"
		
)


// CriarUsuario = Chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
			return
	}


	url := fmt.Spri
	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
			return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
			return
	}

	respostas.JSON(w, response.StatusCode, nil)
		return
}
