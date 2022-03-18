package managerDatasource

import (
	"test/aulaInterface/interfaces"
)

// RefList
//
// Português:
//
// Recebe todos os ponteiros de datasource.
type RefList struct {
	User     interfaces.InterfaceUser     `json:"-"`
	Password interfaces.InterfacePassword `json:"-"`
	UniqueID interfaces.InterfaceUID      `json:"-"`
	Jwt      interfaces.InterfaceJWT      `json:"-"`
}
