package database

import (
	"api/src/config"
	"database/sql"
	//Driver
)

// Open connection wuth database
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionDataBase)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
