package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
)

//createTableMenu (Português): Cria a tabela do menu
func (e *SQLiteUser) createTableUser() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	user (
				id TEXT PRIMARY KEY,
				admin INTEGER,          -- 0: normal user; 1 admin user
				name TEXT,              -- complete name
				nickName TEXT,          -- nick name
				email TEXT,             -- e-mail
				password TEXT           -- password
			);
		`,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = statement.Exec()
	if err != nil {
		util.TraceToLog()
		return
	}
	return
}
