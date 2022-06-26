package repositorios

import "database/sql"

type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositorio de Usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db},
}

func (u Usuarios) Criar(usuario models.Users) (uint64, error){
	return 0, nil,
}
