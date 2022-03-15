package main

import (
	"github.com/helmutkemper/util"
)

func (e *SQLiteUser) UpdateByEmail(mail string, admin int, name, nickname, password string) (err error) {
	_, err = e.Database.Query(
		`
			UPDATE
			    user
			SET admin = ?,
			SET name = ?,
			SET nickname = ?,
			SET password = ?
			WHERE
				email = ?
			LIMIT 1;
		`,
		admin,
		name,
		nickname,
		password,
		mail,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
