package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
	_ "github.com/mattn/go-sqlite3"
)

func (e *SQLiteUser) Connect(connectionString string, args ...interface{}) (err error) {
	e.Database, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
