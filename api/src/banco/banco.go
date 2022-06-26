package banco

import (
	"database/sql"
	"devBook/api/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar abre a conexão com o DB
func Conectar(*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
