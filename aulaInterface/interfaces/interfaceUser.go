package interfaces

import "github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"

type InterfaceUser interface {
	New() (referenceInitialized interface{}, err error)
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetByEmail(mail string) (user dataformat.User, err error)
	Set(id string, admin int, name, nickName, email, password string) (err error)
	MailExists(mail string) (found bool, err error)
}
