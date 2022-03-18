package user

import (
	"errors"
	"log"
	systemDatasource "test/aulaInterface/businessRules/system/managerDatasource"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
)

// getPasswordByEmail
//
// Português:
//
// Regra de negócio para pegar o hash do password do usuário no banco de dados.
//   Entrada:
//     mail: e-mail do usuário.
//   Saída:
//     password: hash do password do usuário;
//     err: objeto de erro padrão do go.
func (e *BusinessRules) getPasswordByEmail(mail string) (password string, err error) {
	var userFromDatasource dataformat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	e.DataSourceUser = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSourceUser.GetByEmail(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	password = userFromDatasource.Password
	return
}
