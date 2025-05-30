package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
	func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
		usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
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
	
		publicacao.AutorID = usuarioID

		if erro = publicacao.Preparar(); erro != nil {
			if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil { 
				respostas.Erro(w, http.StatusBadRequest, erro) 
				return
			}

		}


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

	// BuscarPublicacao = Adiciona uma nova publicação no banco de dados
	func BuscarPublicacao(w http.ResponseWriter, r *http.Request {
		parametros := mux.Vars()
		publicacaoID, erro := st

	}
	
	// // BuscarPublicacoes = Busca as publicações que apareceriam no feed do usuário
	// func BuscarPublicacoes(w http.ResponseWriter, r *http.Request {

	// }	
		
	// // AtualizarPublicacao = Altera os dados de uma publicação
	// func AtualizarPublicacao(w http.ResponseWriter, r *http.Request {

	// }

	// // DeletarPublicacao = Exclui os dados de um publicação
	// func DeletarPublicacao(w http.ResponseWriter, r *http.Request {

	// }	