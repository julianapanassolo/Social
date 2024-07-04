package controllers

import (
	"api/src/modelos"
	"net/http"
	"io/ioutil"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
	func CriarPublicacao(w http.ResponseWriter, r *http.Request {
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

		publicacao.AutorID = usuárioID

		db. erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w. htttp.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
		

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacao.ID, erro := repositorio.Criar(publicacao)
		if erro != nil {
			respostas.Erro(w. htttp.StatusInternalServerError, erro)
			return

		}

		respostas.JSON(w, http.StatusInternalServerError, publicacao)
		
	}

	// // BuscarPublicacao = Adiciona uma nova publicação no banco de dados
	// func BuscarPublicacao(w http.ResponseWriter, r *http.Request {

	// }
	
	// // BuscarPublicacoes = Busca as publicações que apareceriam no feed do usuário
	// func BuscarPublicacoes(w http.ResponseWriter, r *http.Request {

	// }	
		
	// // AtualizarPublicacao = Altera os dados de uma publicação
	// func AtualizarPublicacao(w http.ResponseWriter, r *http.Request {

	// }

	// // DeletarPublicacao = Exclui os dados de um publicação
	// func DeletarPublicacao(w http.ResponseWriter, r *http.Request {

	// }	