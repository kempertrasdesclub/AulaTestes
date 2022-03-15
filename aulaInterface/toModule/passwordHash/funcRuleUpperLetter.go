package passwordHash

import (
	"bytes"
	"errors"
	"log"
)

func (e *Password) ruleUpperLetter(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("A"), []byte("B"), []byte("C"), []byte("D"), []byte("E"), []byte("F"),
		[]byte("G"), []byte("H"), []byte("I"), []byte("J"), []byte("K"), []byte("L"), []byte("M"), []byte("N"),
		[]byte("O"), []byte("P"), []byte("Q"), []byte("R"), []byte("S"), []byte("T"), []byte("V"), []byte("W"),
		[]byte("X"), []byte("Y"), []byte("Z")}
	for _, char = range specialChars {
		if bytes.Contains(password, char) {
			return
		}
	}

	err = errors.New("the password must be one upper case char")
	log.Printf("passwordHash.ruleUpperLetter().error: %v", err.Error())
	return
}
