package main

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/constants"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBUser) verifyInstallUser() (installed bool, err error) {
	var cursor *mongo.Cursor
	var users []dataformat.User
	var user = dataformat.User{Id: constants.KMainUserID, Mail: constants.KMainUserMail}

	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	cursor, err = e.ClientUser.Find(e.Ctx, user.GetIdAndMailAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &users)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(users) > 0
	return
}
