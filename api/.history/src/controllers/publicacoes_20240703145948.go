package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
	func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
		usuárioID, erro := autenticacao.ExtrairUsuarioID(r)
		if erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		corpoRequisicao, erro := ioutil.ReadAll(r.Body)
		if erro != nil {
			respostas.Erro(w, http.StatusUnprocessableEntity, erro)
			return
		}

		var publicacao modelos.Publicacao
		if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil { 
		respostas.Erro(w, http.StatusBadRequest, erro) 
		return
	}

		publicacao.AutorID = usuárioID

		if erro = publicacao.Preparar(); erro 



		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
		

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacao.ID, erro = repositorio.Criar(publicacao)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return

		}

		respostas.JSON(w, http.StatusCreated, publicacao)
		
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