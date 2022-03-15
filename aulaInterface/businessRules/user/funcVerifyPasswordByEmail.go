package user

import (
	"encoding/base64"
	"errors"
	"log"
	"test/aulaInterface/constants"
)

func (e *BusinessRules) VerifyPasswordByEmail(mail, password string) (match bool, err error) {
	var hash []byte
	var passwordFromDatasource string
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.VerifyPasswordByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.VerifyPasswordByEmail().error: %v", err.Error())
		return
	}

	passwordFromDatasource, err = e.getPasswordByEmail(mail)
	if err != nil {
		log.Printf("user.VerifyPasswordByEmail().error: %v", err.Error())
		return
	}

	hash, err = base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(passwordFromDatasource)
	if err != nil {
		log.Printf("user.VerifyPasswordByEmail().error: %v", err.Error())
		return
	}

	match = e.Password.CheckHash([]byte(password), hash)
	return
}
