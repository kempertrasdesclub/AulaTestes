package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

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
