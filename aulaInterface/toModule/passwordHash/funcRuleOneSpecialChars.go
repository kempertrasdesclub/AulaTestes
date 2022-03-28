package passwordHash

import (
	"bytes"
	"errors"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/constants"
	"log"
)

func (e *Password) ruleOneSpecialChars(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("`"), []byte("~"), []byte("!"), []byte("@"), []byte("#"), []byte("$"),
		[]byte("%"), []byte("^"), []byte("&"), []byte("*"), []byte("("), []byte(")"), []byte("-"), []byte("_"),
		[]byte("+"), []byte("="), []byte("["), []byte("{"), []byte("]"), []byte("}"), []byte("|"), []byte("\\"),
		[]byte(";"), []byte(":"), []byte("\""), []byte("'"), []byte("<"), []byte(">"), []byte(","), []byte("."),
		[]byte("/"), []byte("?")}
	for _, char = range specialChars {
		if bytes.Contains(password, char) {
			return
		}
	}

	err = errors.New(constants.KErrorPasswordOneSpecialChar)
	log.Printf("passwordHash.ruleOneSpecialChars().error: %v", err.Error())
	return
}
