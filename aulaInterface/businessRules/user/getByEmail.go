package user

import (
	"errors"
	"log"
	systemDatasource "test/aulaInterface/businessRules/system/managerDatasource"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
	"test/aulaInterface/view/viewUser"
)

// GetByEmail
//
// Português:
//
// Regra de negócio de como carregar dados do usuário por e-mail.
//   Entrada:
//     mail: endereço de e-mail do usuário.
//   Saída:
//     length: quantidade de usuários;
//     user: viewUser.User populada;
//     err: objeto de erro padrão do go.
//
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

	e.DataSourceUser = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSourceUser.GetByEmail(mail)
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
