package main

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/constants"
)

func (e *MongoDBUser) populateInitialUser() (err error) {

	err = e.Set(
		constants.KMainUserID,
		constants.KmainMenuUserAdmin,
		constants.KMainUserName,
		constants.KMainUserNickName,
		constants.KMainUserMail,
		constants.KMainUserPassword,
	)
	if err != nil {
		util.TraceToLog()
		return
	}
	return
}
