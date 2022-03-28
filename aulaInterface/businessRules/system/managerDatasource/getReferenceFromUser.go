package managerDatasource

import (
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/interfaces"
)

// GetReferenceFromUser
//
// Português:
//
// Retorna o datasource do usuário.
//   Saída:
//     datasource: Referência do objeto `User` usado como fonte de dados.
func (e *RefList) GetReferenceFromUser() (datasource interfaces.InterfaceUser) {
	return e.User
}
