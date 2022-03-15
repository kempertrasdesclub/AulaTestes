package viewUser

import (
	"test/aulaInterface/dataformat"
)

func (e *User) Parser(user *dataformat.User) {
	*e = User(*user)
	e.Password = ""
}
