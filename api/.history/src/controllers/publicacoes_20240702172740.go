package controllers

import (
	"api/src/modelos"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
	func CriarPublicacao(w http.ReponserWriter, r *http.Request {
		usuárioID, erro := auteticacao.ExtrairUsuarioID(r)
		if erro != nil {
			respostas.Erro(w. htttp.StatusUnauthorized, erro)
			return
		}

		corpoRequisicao, erro := ioutil.ReadAll(r.Body)
		if erro != nil {
			respostas.Erro(w. htttp.StatusUnprocessableEntity, erro)
			return
		}

		var publicacao modelos.Publicacao
		if erro != nil {
			respostas.Erro(w. htttp.StatusBadRequest, erro)
			return
		}

		
	}

	// BuscarPublicacao = Adiciona uma nova publicação no banco de dados
	func BuscarPublicacao(w http.ReponserWriter, r *http.Request {

	}
	
	// BuscarPublicacoes = Busca as publicações que apareceriam no feed do usuário
	func BuscarPublicacoes(w http.ReponserWriter, r *http.Request {

	}	
		
	// AtualizarPublicacao = Altera os dados de uma publicação
	func AtualizarPublicacao(w http.ReponserWriter, r *http.Request {

	}

	// DeletarPublicacao = Exclui os dados de um publicação
	func DeletarPublicacao(w http.ReponserWriter, r *http.Request {

	}	