package user

import "test/aulaInterface/interfaces"

type BusinessRules struct {
	UniqueID   interfaces.InterfaceUID      `json:"-"`
	Password   interfaces.InterfacePassword `json:"-"`
	DataSource interfaces.InterfaceUser     `json:"-"`
}
