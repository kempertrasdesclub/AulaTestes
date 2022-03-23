package dataTest

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
)

type DataTest struct {
	Id   string
	Name string
}

func (e *DataTest) GetID() (ID interface{}, err error) {
	ID = e.Id
	return
}

func (e *DataTest) getNextId() (id string) {
	return gofakeit.UUID()
}

func (e *DataTest) Get() (data interface{}) {
	return *e
}

func (e *DataTest) Populate() (err error) {
	e.Id = e.getNextId()
	e.Name = gofakeit.Name()
	return
}

func (e *DataTest) Update() (err error) {
	e.Name = gofakeit.Name()
	return
}

func (e *DataTest) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, e)
	if err != nil {
		return
	}

	return
}

func (e *DataTest) MarshalJSON() (data []byte, err error) {
	return json.Marshal(&struct {
		Id   string
		Name string
	}{
		Id:   e.Id,
		Name: e.Name,
	})
}
