package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var User = MongoDBUser{}

type MongoDBUser struct {
	Client     *mongo.Client
	Ctx        context.Context
	CancalFunc context.CancelFunc
	ClientUser *mongo.Collection
}
