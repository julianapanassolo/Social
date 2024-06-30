package repositorios

import (
	"api/src/modelos"	
	"database/sql"
	"fmt"
)

// Usuarios: Representa um repositório de usuários
type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}

// NovoRepositorioDeUsuarios = Cria um repositório de usuários 
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
    return &usuarios{db: db}
}


// Criar = Insere um usuário no banco de dados
func (repositorio *usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",	
	)
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro  // Valor 0 se refere ao "uint64"
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar = Irá trazer todos os usuários que atendem um filtro de nome ou nick
func (repositorio *usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	query := `
		SELECT
			id,
			nome,
			nick,
			email,
			criadoEm
		FROM usuarios WHERE nome LIKE ? OR nick LIKE ?`
	linhas, erro := repositorio.db.Query(query, nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorID = Traz apenas 1 usuário do banco de dados
func (repositorio *usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	 erro := repositorio.db.QueryRow(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", 
		ID, 
	).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.Email,
		&usuario.CriadoEm,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	
	return usuario, nil	
}

// Atualizar = Altera as informações de um usuário no banco de dados
func (repositorio *usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}

// Deletar = Exclui as informações do usuário do banco de dados
func (repositorio *usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio *usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha := repositorio.db.QueryRow("SELECT id, senha FROM usuarios WHERE email = ?", email)

	
	var usuario modelos.Usuario
	// Sacan diretamente na linha retornada	
	if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
		// Verificar se o erro é por não encontrar nenhuma linha
		if erro == sql.ErrNoRows {
			return modelos.Usuario{}, nil  // Se for nenhum usuário encontrado, ele retorna vazio
		}
		return modelos.Usuario{}, erro // Outro erro
	}
	return usuario, nil
 }
// Seguir = Permite que um usuário siga o outro
 func (repositório Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositório.db.Prepare(
		"INSERT INTO seguidores (usuario_id)"
	)

 }
