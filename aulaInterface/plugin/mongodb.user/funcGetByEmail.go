package main

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/constants"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBUser) GetByEmail(mail string) (user dataformat.User, err error) {
	var cursor *mongo.Cursor
	var userSlice []dataformat.User
	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	user = dataformat.User{Mail: mail}
	cursor, err = e.ClientUser.Find(e.Ctx, user.GetMailAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &userSlice)
	if err != nil {
		util.TraceToLog()
		return
	}

	if len(userSlice) != 0 {
		user = userSlice[0]
	}

	return
}
