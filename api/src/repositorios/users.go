package repositorios

import (
	"database/sql"
	"devBook/api/src/models"
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
