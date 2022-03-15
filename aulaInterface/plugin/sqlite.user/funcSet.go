package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
)

// Set (Português): Adiciona um novo usuário
func (e *SQLiteUser) Set(id string, admin int, name, nickName, email, password string) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.user (id, admin, name, nickName, email, password) VALUES(?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = statement.Exec(id, admin, name, nickName, email, password)
	if err != nil {
		util.TraceToLog()
		return
	}
	return
}
