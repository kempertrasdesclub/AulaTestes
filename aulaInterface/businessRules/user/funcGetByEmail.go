package user

import (
	"errors"
	"log"
	systemDatasource "test/aulaInterface/businessRules/system/datasource"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
	"test/aulaInterface/view/viewUser"
)

// GetByEmail (Português):
func (e *BusinessRules) GetByEmail(mail string) (length int, user viewUser.User, err error) {
	var userFromDatasource dataformat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	if userFromDatasource.Mail != "" {
		length = 1
	}

	user = viewUser.User{}
	user.Parser(&userFromDatasource)

	return
}
