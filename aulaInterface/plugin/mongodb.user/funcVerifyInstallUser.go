package main

import (
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/mongo"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
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
