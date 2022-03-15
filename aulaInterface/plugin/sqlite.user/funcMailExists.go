package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
)

func (e *SQLiteUser) MailExists(mail string) (found bool, err error) {
	var rows *sql.Rows
	rows, err = e.Database.Query(
		`
			SELECT 
				id 
			FROM 
				user 
			WHERE 
				email = ?
		`,
		mail,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	found = rows.Next()
	return
}
