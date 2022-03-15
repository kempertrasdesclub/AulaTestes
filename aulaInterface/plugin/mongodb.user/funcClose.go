package main

func (e *MongoDBUser) Close() (err error) {
	err = e.Client.Disconnect(e.Ctx)
	return
}
