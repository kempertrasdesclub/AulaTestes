package viewUser

import (
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

func (e *User) Parser(user *dataformat.User) {
	*e = User(*user)
	e.Password = ""
}
