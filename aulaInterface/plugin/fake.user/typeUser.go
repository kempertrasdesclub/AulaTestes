package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

var User = FakeUser{}

type FakeUser struct{}

func (e *FakeUser) Set(id string, admin int, name, nickName, email, password string) (err error) {
	return
}

func (e *FakeUser) New() (referenceInitialized interface{}, err error) {
	return e, nil
}

func (e *FakeUser) MailExists(mail string) (found bool, err error) {
	found = true
	return
}

func (e *FakeUser) Install() (err error) {
	return
}

func (e *FakeUser) GetByEmail(mail string) (user dataformat.User, err error) {
	var nameFirst = gofakeit.FirstName()
	var nameLast = gofakeit.LastName()
	user = dataformat.User{
		Id:       gofakeit.UUID(),
		Admin:    gofakeit.RandomInt([]int{0, 1}),
		Name:     nameFirst + " " + nameLast,
		NickName: nameFirst + "." + nameLast,
		Mail:     nameFirst + "." + nameLast + "@company.com",
		Password: "**********",
	}

	return
}

func (e *FakeUser) Connect(connectionString string, args ...interface{}) (err error) {
	return
}

func (e *FakeUser) Close() (err error) {
	return
}
