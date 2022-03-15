package main

import (
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
)

func (e *MongoDBUser) GetAll() (users []dataformat.User, length int, err error) {
	var cursor *mongo.Cursor

	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)
	cursor, err = e.ClientUser.Find(e.Ctx, bson.M{})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &users)
	if err != nil {
		util.TraceToLog()
		return
	}

	length = len(users)
	return
}
