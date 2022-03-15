package main

import (
	"context"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (e *MongoDBUser) Connect(connectionString string, args ...interface{}) (err error) {
	e.Client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		util.TraceToLog()
		return
	}

	e.Ctx = context.Background()
	err = e.Client.Connect(e.Ctx)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.Client.Ping(e.Ctx, readpref.Primary())
	if err != nil {
		util.TraceToLog()
	}

	return
}
