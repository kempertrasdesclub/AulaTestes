package passwordHash

import "golang.org/x/crypto/bcrypt"

func (e *Password) CheckHash(password, hash []byte) (match bool) {
	var err = bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
