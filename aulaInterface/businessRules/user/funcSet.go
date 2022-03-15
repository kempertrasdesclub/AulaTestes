package user

import (
	"encoding/base64"
	"errors"
	"log"
	systemDatasource "test/aulaInterface/businessRules/system/datasource"
	"test/aulaInterface/constants"
)

func (e *BusinessRules) Set(idMenu string, admin int, name, nickName, mail, password string) (err error) {
	var matched bool
	var hash []byte

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.SetMainMenu().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.SetMainMenu().error: %v", err.Error())
		return
	}

	hash, err = e.Password.MakeHash([]byte(password))
	if err != nil {
		log.Printf("user.SetMainMenu().error: %v", err.Error())
		return
	}

	password = base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(hash)

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	err = e.DataSource.Set(e.UniqueID.Get(), admin, name, nickName, mail, password)
	return
}
