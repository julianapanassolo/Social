package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao = Chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r*http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string {
		"titulo": r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("$/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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
}

// CurtirPublicacao = Chama a API para curtir uma publicação
func CurtirPublicacao(w http.ResponseWriter, r*http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacao"])

}