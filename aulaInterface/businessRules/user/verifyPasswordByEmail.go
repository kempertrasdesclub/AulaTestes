package user

import (
	"encoding/base64"
	"errors"
	"log"
	"test/aulaInterface/constants"
)

// VerifyPasswordByEmail
//
// Português:
//
// Verifica se a senha recebida pelo frontend confere com o hash da senha salva na fonte de dados.
//   Entrada:
//     mail: e-mail do usuário;
//     password: senha do usuário.
//   Saída:
//     match: true quando a senha confere;
//     err: objeto de erro padrão do go.
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
