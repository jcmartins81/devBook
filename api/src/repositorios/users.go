package repositorios

import (
	"database/sql"
	"devBook/api/src/models"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositorio de UsuÃ¡rios
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

	linhas, erro := repositorio.db.Query("select id, name, nick, email, createdAT, from usuarios where name LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick)

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
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarPorID traz um usuÃ¡tio do banco de dados
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

func (repositorio Usuarios) AtualizarUsuario(ID uint64, usuario models.User) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Name, usuario.NickName, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}
//Deletar exclui as informações de um usuário no BAnco de Dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuario por email e retorna seu id e senha
func (repositorio Usuarios) BuscarPorEmail(email string) (models.User, error) {
	linha, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?", email
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer linha.Close()

	var usuario models.User

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.User{}, erro
		}
	}

	return usuario, nil
}
