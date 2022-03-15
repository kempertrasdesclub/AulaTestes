package main

import (
	"test/aulaInterface/constants"
)

func (e *MongoDBUser) createTableUser() (err error) {
	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)
	return
}
