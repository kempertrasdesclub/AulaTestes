package main

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/constants"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

// GetByEmail (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteUser) GetByEmail(mail string) (user dataformat.User, err error) {
	var rows *sql.Rows
	rows, err = e.Database.Query(
		`
			SELECT
				id,
				admin,
				name,
				nickName,
				email,
				password
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

	var id string
	var admin int
	var name string
	var nickName string
	var password string

	if rows.Next() {
		err = rows.Scan(&id, &admin, &name, &nickName, &mail, &password)
		if err != nil {
			util.TraceToLog()
			return
		}

		user.Id = id
		user.Admin = admin
		user.Name = name
		user.NickName = nickName
		user.Mail = mail
		user.Password = password
	} else {
		err = errors.New(constants.KErrorUserNotFound)
		util.TraceToLog()
	}

	return
}
