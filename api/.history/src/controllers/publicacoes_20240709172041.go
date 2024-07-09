package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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
		respostas.Erro(w, http.StatusBadRequest, erro) 
			return
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

	// BuscarPublicacao = Adiciona uma única publicação no banco de dados
	func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
		parametros := mux.Vars(r)
		publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
		if erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}

		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacao, erro := repositorio.BuscarPorID(publicacaoID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return

		}
		respostas.JSON(w, http.StatusOK, publicacao)

	}
	
	// BuscarPublicacoes = Busca as publicações que apareceriam no feed do usuário
	func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
		usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
		if erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
	
		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
	
		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacoes, erro := repositorio.Buscar(usuarioID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	
		respostas.JSON(w, http.StatusOK, publicacoes)
	}
		
	// AtualizarPublicacao = Altera os dados de uma publicação
	func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
		usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
		if erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		parametros := mux.Vars(r)
		publicacaoID, erro := strconv.ParseUint(parametros["puplicacaoId"], 10, 64)
		if erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}
		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		if publicacaoSalvaNoBanco.AutorID != usuarioID {
			respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
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

		if erro = publicacao.Preparar(); erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}

		if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		respostas.JSON(w, http.StatusNoContent, nil)	
}

	// DeletarPublicacao = Excluiu os dados de uma publicação
	func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
		usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
		if erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		parametros := mux.Vars(r)
		publicacaoID, erro := strconv.ParseUint(parametros["puplicacaoId"], 10, 64)
		if erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}
		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		if publicacaoSalvaNoBanco.AutorID != usuarioID {
			respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua"))
			return
		}

		if erro = repositorio.Deletar(publicacaoID); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		respostas.JSON(w, http.StatusNoContent, nil)
	} 

	// BuscarPublicacoesPorUsuario = Busca por publicaões de um único usuário dentro do banco da dados
	func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
		parametros := mux.Vars(r)
		usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
		if erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}

		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		respostas.JSON(w, http.StatusOK, publicacoes)

	}

	// CurtirPublicacao =  Irá adicionar uma curtida na publicação
	func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
		parametros := mux.Vars(r)
		publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
		if erro != nil {
			respostas.Erro(w, http.StatusBadRequest, erro)
			return
		}

		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)
		if erro = repositorio.Curtir(publicacaoID); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		respostas.JSON(w, http.StatusNoContent, nil )

	}

	// DescurtirPublicacao =  Irá tirar uma curtida na publicação
	func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
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
		respostas.Erro(w, http.StatusBadRequest, erro) 
			return
		}
		

		db, erro := banco.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
		

		repositorio := repositorios.NovoRepositorioDePublicacoes(db)		
		if erro = repositorio.Descurtir()
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return

		}

		respostas.JSON(w, http.StatusCreated, publicacao)

		
	}