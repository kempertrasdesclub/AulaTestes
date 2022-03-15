package datasource

import (
	"test/aulaInterface/interfaces"
)

// GetReferenceFromUser (Português): Retorna o datasource do usuário
func (e *RefList) GetReferenceFromUser() (datasource interfaces.InterfaceUser) {
	return e.User
}
