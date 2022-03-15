package main

import (
	"github.com/helmutkemper/util"
	"test/aulaInterface/constants"
)

func (e *MongoDBUser) New() (referenceInitialized interface{}, err error) {
	err = e.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return e, nil
}
