package passwordHash

import (
	"log"
)

func (e *Password) newPasswordRules(password []byte) (err error) {
	err = e.ruleLength(password)
	if err != nil {
		log.Printf("passwordHash.newPasswordRules().error: %v", err.Error())
		return
	}

	err = e.ruleOneSpecialChars(password)
	if err != nil {
		log.Printf("passwordHash.newPasswordRules().error: %v", err.Error())
		return
	}

	err = e.ruleUpperLetter(password)
	if err != nil {
		log.Printf("passwordHash.newPasswordRules().error: %v", err.Error())
		return
	}

	err = e.ruleLowerCase(password)
	if err != nil {
		log.Printf("passwordHash.newPasswordRules().error: %v", err.Error())
	}
	return
}
