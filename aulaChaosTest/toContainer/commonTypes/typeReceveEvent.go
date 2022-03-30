package commonTypes

type ReceiveEvent int

const (
	KCreate ReceiveEvent = iota + 1
	//KRead
	KUpdate
	KDelete
)
