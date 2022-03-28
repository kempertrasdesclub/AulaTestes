package main

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

func (e *MongoDBUser) Set(id string, admin int, name, nickName, email, password string) (err error) {

	_, err = e.ClientUser.InsertOne(
		e.Ctx,
		dataformat.User{
			Id:       id,
			Admin:    admin,
			Name:     name,
			NickName: nickName,
			Mail:     email,
			Password: password,
		},
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
