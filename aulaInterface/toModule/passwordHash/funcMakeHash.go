package passwordHash

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (e *Password) MakeHash(password []byte) (hash []byte, err error) {
	err = e.newPasswordRules(password)
	if err != nil {
		log.Printf("passwordHash.MakeHash().error: %v", err.Error())
		return
	}

	hash, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("passwordHash.MakeHash().error: %v", err.Error())
	}
	return
}
