package main

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

func (e *MongoDBUser) MailExists(mail string) (found bool, err error) {
	var user dataformat.User
	user, err = e.GetByEmail(mail)
	if err != nil {
		util.TraceToLog()
		return
	}

	found = user.Mail != ""
	return
}
