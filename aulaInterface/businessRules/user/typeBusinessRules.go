package user

import "test/aulaInterface/interfaces"

// BusinessRules
//
// Português:
//
// Recebe os ponteiros para todos os objetos usados pelas regras de negócios.
type BusinessRules struct {
	// UniqueID
	//
	// Português:
	//
	// Objeto gerador do ID Único.
	UniqueID interfaces.InterfaceUID `json:"-"`

	// Password
	//
	// Português:
	//
	// Objeto responsável pelo hash da senha do usuário.
	Password interfaces.InterfacePassword `json:"-"`

	// DataSourceUser
	//
	// Português:
	//
	// Objeto responsável pela fonte de dados referente ao usuário.
	DataSourceUser interfaces.InterfaceUser `json:"-"`
}
