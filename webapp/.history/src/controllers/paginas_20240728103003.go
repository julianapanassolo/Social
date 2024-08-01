package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// CarregarTelaDeLogin = Irá realizar a renderização da página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)

}

// CarregarPaginaDeCadastroDeUsuario = Irá carregar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal = Carrega a página principal com as publicações - home
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas
	}

	utils.ExecutarTemplate(w, "home.html", publicacoes)

}
