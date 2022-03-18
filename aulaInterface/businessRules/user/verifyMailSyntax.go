package user

import (
	"log"
	"regexp"
)

// verifyMailSyntax
//
// Português:
//
// Verifica se o e-mail é um formato válido pelo RFC 5322 Official Standard
//   Entrada:
//     mail: e-mail do usuário.
//   Saída:
//     match: true para formato de e-mail correto.
//     err: objeto de erro padrão do go.
func (e *BusinessRules) verifyMailSyntax(mail string) (match bool, err error) {
	match, err = regexp.MatchString(kRFC5322OfficialStandard, mail)
	if err != nil {
		log.Printf("user.verifyMailSyntax().error: %v", err.Error())
	}

	return
}
