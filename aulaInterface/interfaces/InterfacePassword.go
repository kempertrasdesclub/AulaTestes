package interfaces

type InterfacePassword interface {
	MakeHash(password []byte) (hash []byte, err error)
	CheckHash(password, hash []byte) (match bool)
}
