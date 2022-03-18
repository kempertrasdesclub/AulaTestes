package user

import (
	"encoding/base64"
	"errors"
	"log"
	systemDatasource "test/aulaInterface/businessRules/system/managerDatasource"
	"test/aulaInterface/constants"
)

// Set
//
// Português:
//
// Adiciona um novo usuário.
//   Entrada:
//     admin: 0 para user, ou 1 para administrador;
//     name: Nome completo do usuário;
//     nickName: Apelido do usuário apresentado no frontend;
//     mail: E-mail do usuário;
//     password: Senha do usuário.
//   Saída:
//     err: Objeto de erro padrão do go.
func (e *BusinessRules) Set(admin int, name, nickName, mail, password string) (err error) {
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

	e.DataSourceUser = systemDatasource.Linker.GetReferenceFromUser()
	err = e.DataSourceUser.Set(e.UniqueID.Get(), admin, name, nickName, mail, password)
	return
}
