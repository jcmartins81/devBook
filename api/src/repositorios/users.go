package repositorios

import (
	"database/sql"
	"devBook/api/src/models"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositorio de Usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.User) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Name, usuario.NickName, usuario.Email, usuario.Password)

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.User, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, createdAT, from usuarios where name LIKE ?
		or nick LIKE ?", nomeOuNick, nomeOuNick
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.User

	for linhas.Next() {
		var usuario models.User

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.NickName,
			&usuario.Email,
			&usuario.CreatedAt,
		); erro != nil{
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
	
}

// BuscarPorID traz um usuátio do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (models.User, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, NickName, email, criatedAt from usuarios where id = ?", 
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer linhas.Close()
	
	var usuario models.User

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.NickName,
			&usuario.Email,
			&usuario.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return usuario, nil
}

